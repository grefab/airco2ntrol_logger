<template>
    <div id="app">
        <MainScreen msg="Welcome to Your Vue.js App"/>
    </div>
</template>

<script>
    import MainScreen from './components/MainScreen.vue'

    export default {
        name: 'app',
        components: {
            MainScreen
        }
    }
</script>

<style>
</style>
