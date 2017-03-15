<template>
    <div class="home">
        <h3>Contix Editor</h3>
        <div class="line"></div>

        <div id="pick-file">
            <el-row :gutter="5">
                <el-col :span="22">
                    <input type="file" class="load-configuire-file" v-on:change="loadConfiguireFile" accept="application/x-yaml, .yaml, .yml">
                    <el-input v-bind:placeholder="browseInputBoxPlaceholder" readonly>
                        <el-button slot="append" v-on:click="browse">Browse</el-button>
                    </el-input>
                </el-col>
                <el-col :span="2" class="text-center full-width-button">
                    <el-button v-bind:disabled="isLoaded === false" v-on:click="saveConfigureFile">Save</el-button>
                </el-col>
            </el-row>
        </div>
        <div class="line"></div>

        <div class="editor" v-if="isLoaded === true">
            <el-tabs type="border-card" v-model="editorCurrentTab" @tab-click="clickEditorTab">
                <el-tab-pane label="Performances" name="performances">
                    <home-performances v-bind:configs="configs"></home-performances>
                </el-tab-pane>
                <el-tab-pane label="Mail" name="mail">
                    <home-mail v-bind:configs="configs"></home-mail>
                </el-tab-pane>
                <el-tab-pane label="User Agents" name="user-agents">
                    User Agents
                </el-tab-pane>
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
import HomePerformances from './home/Performances.vue'
import HomeMail from './home/Mail.vue'

export default {
    name: 'home',

    components: {
        'home-performances': HomePerformances,
        'home-mail': HomeMail,
    },

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

            const reader = new FileReader()

            reader.onload = () => {
                const configs = yaml.safeLoad(reader.result)

                const cond1 = configs.hasOwnProperty("performances")
                const cond2 = configs.hasOwnProperty("mail")
                const cond3 = configs.hasOwnProperty("user_agents")

                if (cond1 === true && cond2 === true && cond3 === true) {
                    this.configs  = configs
                    this.isLoaded = true
                }else{
                    this.configs  = configs
                    this.file     = {}
                    this.isLoaded = true

                    alert("Invalid file format")
                }

                // Reset input[type=file] to fix only load once problem
                document.querySelector(".load-configuire-file").value = ""
            }

            if (event.target.files.length > 0) {
                this.file = event.target.files[0]

                reader.onload.bind(this)
                reader.readAsText(this.file)
            }
        },

        saveConfigureFile() {
            const dump = yaml.safeDump(this.configs)
            const url  = "data:application/x-yaml;base64," + this.base64EncodeUnicode(dump)
            const aTag = document.createElement("a")

            aTag.style    = "display: none"
            aTag.href     = url
            aTag.download = "cron-task.yaml"

            document.body.appendChild(aTag)

            aTag.click()

            document.body.removeChild(aTag)
        },

        clickEditorTab(tab, event) {

        },

        base64EncodeUnicode(text) {
            return btoa(encodeURIComponent(text).replace(/%([0-9A-F]{2})/g, (match, p1) => {
                return String.fromCharCode('0x' + p1);
            }))
        },
    }
}
</script>
