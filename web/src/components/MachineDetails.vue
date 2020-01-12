<template>
    <div>
        <h3 style="text-align: center;">{{ machineName }}</h3>
        <b-row>
            <b-col cols="12">
                <b-card>
                    <h5 class="card-title">Rejection History
                        <small class="card-subtitle text-muted">rejection rate in % for the past 24h</small>
                    </h5>
                    <div id="history-chart"></div>
                    <vue-c3 :handler="rejectionHistoryHandler" style="height: 160px;"/>
                </b-card>
            </b-col>
        </b-row>
        <b-row>
            <b-col cols="12">
                <b-card>
                    <h5 class="card-title">Defect Details
                        <small class="card-subtitle text-muted">
                            contribution per parameter to rejections of the last hour
                        </small>
                    </h5>
                    <div class="bar-chart"></div>
                    <vue-c3 :handler="defectDetailsHandler"/>
                </b-card>
            </b-col>
        </b-row>
    </div>
</template>

<script>
    import {getStorageServiceEndpoint} from "@/helper"
    import {StatisticsClient} from '../proto/isotronic/vialserver/storage/api/proto/StatisticsService_grpc_web_pb'
    import {
        MachineRequest,
        MachineTimeFrameRequest,
        MachineShiftRegisterTimeFrameRequest
    } from '../proto/isotronic/vialserver/storage/api/proto/StatisticsService_pb'
    import {TimeFrame} from '../proto/isotronic/common/proto/TimeFrame_pb'
    import {MachineClient} from '../proto/isotronic/vialserver/storage/api/proto/MachineService_grpc_web_pb'
    import Vue from 'vue'
    import VueC3 from 'vue-c3'
    import 'c3/c3.min.css'

    export default {
        name: 'MachineDetails',
        components: {VueC3},
        props: {
            machineName: String,
        },
        data() {
            return {
                rejectionHistoryHandler: new Vue(),
                defectDetailsHandler: new Vue(),
            }
        },
        created() {
            this.statistics = new StatisticsClient(getStorageServiceEndpoint(), null, null);
            this.machines = new MachineClient(getStorageServiceEndpoint(), null, null);
        },
        mounted() {
            this.setupHistory();
            this.setupParameterPerformance();

            let machineRequest = new MachineRequest();
            machineRequest.setMachineName(this.machineName);
            this.statistics.get_rejection_rate_history(machineRequest, {}, (err, rejectionRateHistory) => {
                let history = rejectionRateHistory.toObject();
                this.updateRejectionHistory(history.historyList);
            });
        },
        methods: {
            setupHistory: function () {
                const options = {
                    data: {
                        type: 'bar',
                        x: 'x',
                        columns: [],
                        labels: true
                    },
                    axis: {
                        x: {
                            type: 'category',
                            tick: {
                                rotate: 90,
                                multiline: false
                            }
                        },
                        y: {
                            min: 0,
                            max: 20,
                            padding: {
                                bottom: 0
                            },
                            tick: {
                                values: [0, 10, 25, 50, 100]
                            }
                        }
                    },
                    grid: {
                        y: {
                            lines: [
                                {value: 10},
                                {value: 25},
                                {value: 50}
                            ]
                        }
                    },
                    legend: {
                        show: false
                    }
                };

                this.rejectionHistoryHandler.$emit('init', options);
            },
            updateRejectionHistory: function (newValues) {
                const categories = ['x'];
                const reject = ['reject %'];

                for (let i = 0; i < newValues.length; ++i) {
                    categories.push(newValues[i].hourOfDay + "-" + (newValues[i].hourOfDay + 1) + "h");
                    reject.push(newValues[i].rejectionRate.toFixed(2));
                }

                this.rejectionHistoryHandler.$emit('dispatch', (chart) => chart.load({
                    columns: [categories, reject]
                }));
            },

            setupParameterPerformance: function () {
                const options = {
                    data: {
                        type: 'bar',
                        x: 'x',
                        columns: [],
                        groups: [
                            ['out<', 'out>']
                        ],
                        order: false,
                        labels: true
                    },
                    axis: {
                        x: {
                            type: 'category',
                            tick: {
                                multiline: false
                            }
                        },
                        rotated: true
                    }
                };
                this.defectDetailsHandler.$emit('init', options);

                const machineName = this.machineName;

                let machineTimeFrameRequest = new MachineTimeFrameRequest();
                machineTimeFrameRequest.setMachineName(machineName);
                let timeFrame = new TimeFrame();
                // TODO(gregor): Select proper time-frame here. If not set, it defaults to the last 60 minutes.
                // timeFrame.setBeginTimestamp("2019-12-10T00:00:00.000Z");
                // timeFrame.setEndTimestamp("2019-12-10T01:00:00.00Z");
                machineTimeFrameRequest.setTimeFrame(timeFrame);
                this.statistics.get_active_shift_registers(machineTimeFrameRequest, {}, (err, activeShiftRegisters) => {
                    activeShiftRegisters.toObject().activeShiftRegistersList.forEach((shiftRegisterName) => {
                        let machineShiftRegisterTimeFrameRequest = new MachineShiftRegisterTimeFrameRequest();
                        machineShiftRegisterTimeFrameRequest.setMachineName(machineName);
                        machineShiftRegisterTimeFrameRequest.setTimeFrame(timeFrame);
                        machineShiftRegisterTimeFrameRequest.setShiftRegisterName(shiftRegisterName);
                        this.statistics.get_parameter_performance(machineShiftRegisterTimeFrameRequest, {}, (err, response) => {
                            let pp = response.toObject();
                            let categories = ['x'];
                            let outSmall = ['out< %'];
                            let outBig = ['out> %'];
                            let out = ['outΣ %'];
                            pp.parameterPerformanceList.forEach((parameter) => {
                                categories.push('(' + parameter.shortDescription + ') ' + parameter.description);
                                const tooSmallPercent = 100 * (parameter.tooSmallCount ? parameter.tooSmallCount : 0) / parameter.totalCount;
                                const tooBigPercent = 100 * (parameter.tooBigCount ? parameter.tooBigCount : 0) / parameter.totalCount;
                                const badPercent = 100 * (parameter.badCount ? parameter.badCount : 0) / parameter.totalCount;
                                outBig.push(tooBigPercent.toFixed(2));
                                outSmall.push(tooSmallPercent.toFixed(2));
                                out.push(badPercent.toFixed(2));
                            });
                            this.defectDetailsHandler.$emit('dispatch', (chart) => chart.load({
                                columns: [categories, outSmall, outBig, out]
                            }));
                        });
                    });
                });
            },
        },
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
