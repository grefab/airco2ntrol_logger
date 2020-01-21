<template>
    <div>
        <plotlygraph :propData="plotInfo" divId="plot" style="height: 600px"/>
    </div>
</template>

<script>
    import 'c3/c3.min.css'
    import moment from 'moment'
    import plotlygraph from './PlotlyGraph'
    import {StorageClient} from '../pb/api_grpc_web_pb'
    import {getStorageServiceEndpoint} from "@/helper"
    import {TimeFrame} from '../pb/api_pb'
    import {Timestamp} from 'google-protobuf/google/protobuf/timestamp_pb'

    export default {
        name: 'Dashboard',
        data() {
            return {
                plotInfo: {
                    data: [{
                        x: [new Date(), 1, 2, 3],
                        y: [4, 1, 2, 1],
                        type: 'scatter'
                    }]
                },
            }
        },
        components: {plotlygraph},
        methods: {},
        created() {
            this.storage = new StorageClient(getStorageServiceEndpoint(), null, null);
        },
        mounted() {
            // prepare request
            let timeFrame = new TimeFrame();
            {
                let now = moment();
                let nowTimestamp = new Timestamp();
                nowTimestamp.setSeconds(now.unix());
                nowTimestamp.setNanos(now.milliseconds() * 1000);
                let then = now.subtract(24, 'hours');
                let thenTimestamp = new Timestamp();
                thenTimestamp.setSeconds(then.unix());
                thenTimestamp.setNanos(then.milliseconds() * 1000);
                timeFrame.setFrom(thenTimestamp);
                timeFrame.setTo(nowTimestamp);
            }

            // execute request
            this.storage.getBatch(timeFrame, {}, (err, batch) => {
                    let dates = [];
                    let co2s = [];
                    let temps = [];
                    batch.getItemsList().forEach((e) => {
                        let ts = new Date(e.getTimestamp().getSeconds() * 1000);
                        dates.push(ts);
                        co2s.push(e.getCo2());
                        temps.push(e.getTmp());
                    });

                    this.plotInfo = {
                        data: [
                            {
                                x: dates,
                                y: co2s,
                                type: 'scatter'
                            },
                            {
                                x: dates,
                                y: temps,
                                yaxis: 'y2',
                                type: 'scatter'
                            },
                        ],
                        layout: {
                            title: 'livingroom air quality',
                            xaxis: {
                                rangeselector: {
                                    buttons: [{
                                        step: 'month',
                                        stepmode: 'backward',
                                        count: 1,
                                        label: '1m'
                                    }, {
                                        step: 'month',
                                        stepmode: 'backward',
                                        count: 6,
                                        label: '6m'
                                    }, {
                                        step: 'year',
                                        stepmode: 'todate',
                                        count: 1,
                                        label: 'YTD'
                                    }, {
                                        step: 'year',
                                        stepmode: 'backward',
                                        count: 1,
                                        label: '1y'
                                    }, {
                                        step: 'all',
                                    }],
                                    rangeslider: {}
                                }
                            },
                            yaxis: {
                                title: 'co2 ppm',
                                fixedrange: true
                            },
                            yaxis2: {
                                title: 'temp °C',
                                overlaying: 'y',
                                side: 'right'
                            }
                        }
                    }
                }
            );
        },
        beforeDestroy() {
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
