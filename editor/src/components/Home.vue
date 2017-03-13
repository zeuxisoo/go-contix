<template>
    <div class="home">
        <h3>Contix Editor</h3>
        <div class="line"></div>

        <div id="pick-file">
            <input type="file" class="load-configuire-file" v-on:change="loadConfiguireFile" accept="application/x-yaml, .yaml, .yml">
            <el-input v-bind:placeholder="browseInputBoxPlaceholder" readonly>
                <el-button slot="append" v-on:click="browse">Browse</el-button>
            </el-input>
        </div>
        <div class="line"></div>

        <div class="editor" v-if="isLoaded === true">
            <el-tabs type="border-card" v-model="editorCurrentTab" @tab-click="clickEditorTab">
                <el-tab-pane label="Performances" name="performances">Performances</el-tab-pane>
                <el-tab-pane label="Mail" name="mail">Mail</el-tab-pane>
                <el-tab-pane label="User Agents" name="user-agents">User Agents</el-tab-pane>
            </el-tabs>
        </div>
    </div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.load-configuire-file {
    display: none;
}
</style>

<script>
import yaml from 'js-yaml'

export default {
    name: 'home',

    data () {
        return {
            configs : {},
            file    : {},
            isLoaded: false,

            editorCurrentTab: 'performances'
        }
    },

    computed: {
        browseInputBoxPlaceholder() {
            if (typeof this.file.name !== "undefined") {
                return this.file.name
            }else{
                return  "Please load the configuire file first"
            }
        }
    },

    methods: {
        browse() {
            const mouseEvents = document.createEvent("MouseEvents")
            mouseEvents.initEvent("click", true, false)

            const loadConfigureFile = document.querySelector(".load-configuire-file")
            loadConfigureFile.dispatchEvent(mouseEvents)
        },

        loadConfiguireFile(event) {
            this.configs = {}
            this.isLoaded = false

            var self = this
            const reader = new FileReader()

            reader.onload = () => {
                const configs = yaml.safeLoad(reader.result)

                const cond1 = configs.hasOwnProperty("performances")
                const cond2 = configs.hasOwnProperty("mail")
                const cond3 = configs.hasOwnProperty("user_agents")

                if (cond1 === true && cond2 === true && cond3 === true) {
                    self.configs  = configs
                    self.isLoaded = true
                }else{
                    self.configs  = configs
                    self.file     = {}
                    self.isLoaded = true

                    alert("Invalid file format")
                }

                // Reset input[type=file] to fix only load once problem
                document.querySelector(".load-configuire-file").value = ""
            };

            if (event.target.files.length > 0) {
                this.file = event.target.files[0]

                reader.readAsText(this.file)
            }
        },

        clickEditorTab(tab, event) {

        }
    }
}
</script>
