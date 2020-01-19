<template>
    <div>
        <b-navbar toggleable="lg" type="light" class="bg-light">
            <b-navbar-brand href="#">kanarienvogel</b-navbar-brand>
            <b-navbar-toggle target="nav-collapse"/>

            <b-collapse id="nav-collapse" is-nav>
                <!--
                <b-navbar-nav>
                    &lt;!&ndash; add more tabs here&ndash;&gt;
                    <b-nav-item
                            href="#"
                            @click="activeApp='dashboard'"
                            v-bind:class="{ active: activeApp === 'dashboard' }">
                        History
                    </b-nav-item>
                </b-navbar-nav>
                -->

                <!-- Right aligned nav items -->
                <b-navbar-nav class="ml-auto">
                    <b-nav-form>
                        <a style="color: #414141" href="https://isotronic.de" target="_blank" rel="noopener noreferrer">
                            <span class="isotronic">ISOTRONIC</span>
                        </a>
                    </b-nav-form>
                </b-navbar-nav>
            </b-collapse>
        </b-navbar>
        <b-container>
            <!-- add more full screen components here-->
            <!-- <Dashboard v-if="activeApp === 'dashboard'"/> -->
            <Dashboard/>
        </b-container>
    </div>
</template>

<script>
    import Dashboard from './Dashboard.vue'

    export default {
        name: 'MainScreen',
        data() {
            return {
                activeApp: String
            }
        },
        components: {
            Dashboard,
        },
        created: function () {
            this.activeApp = "dashboard"
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
    @font-face {
        font-family: "Barlow";
        src: url('/static/Barlow-Bold.ttf') format("truetype");
    }

    .isotronic {
        font-family: "Barlow", sans-serif;
        line-height: 1em;
    }
</style>
