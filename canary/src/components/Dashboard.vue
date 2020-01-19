<template>
    <div>
        <vue-c3 :handler="historyHandler" style="height: 160px;"/>
    </div>
</template>

<script>
    import Vue from 'vue'
    import VueC3 from 'vue-c3'
    import 'c3/c3.min.css'
    import {StorageClient} from '../pb/api_grpc_web_pb'
    import {getStorageServiceEndpoint} from "@/helper"
    import {AirQuality} from '../pb/api_pb'
    import {Timestamp} from 'google-protobuf/google/protobuf/timestamp_pb'

    export default {
        name: 'Dashboard',
        data() {
            return {
                historyHandler: new Vue(),
            }
        },
        components: {VueC3},
        created() {
            this.storage = new StorageClient(getStorageServiceEndpoint(), null, null);
        },
        mounted() {
            this.setupHistory();

            let sinceWhen = new Timestamp();
            let call = this.storage.getSince(sinceWhen, {});
            call.on('data', function (airquality) {
                console.log(airquality);
            });
            call.on('end', function () {
                // The server has finished sending
            });
            call.on('error', function (e) {
                // An error has occurred and the stream has been closed.
            });
            call.on('status', function (status) {
                // process status
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

                this.historyHandler.$emit('init', options);
            }
            ,
            updateRejectionHistory: function (newValues) {
                const categories = ['x'];
                const reject = ['reject %'];

                for (let i = 0; i < newValues.length; ++i) {
                    categories.push(newValues[i].hourOfDay + "-" + (newValues[i].hourOfDay + 1) + "h");
                    reject.push(newValues[i].rejectionRate.toFixed(2));
                }

                this.historyHandler.$emit('dispatch', (chart) => chart.load({
                    columns: [categories, reject]
                }));
            }
            ,
        }
        ,
        beforeDestroy() {
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
