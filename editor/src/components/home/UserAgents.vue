<template>
    <div id="home-user-agents">
        <el-row :gutter="20" class="tab-heading">
            <el-col :span="3">
                <span class="tab-subject">Agents</span>
            </el-col>
            <el-col :span="21" class="text-right">
                <el-button type="success" size="small" class="button-create" v-on:click="showCreateUserAgentForm = true">
                    <i class="el-icon-plus el-icon--left"></i>
                    Create
                </el-button>
            </el-col>
        </el-row>
        <div class="line"></div>

        <el-row :gutter="20">
            <el-col :span="3">Name</el-col>
            <el-col :span="19">Agent</el-col>
            <el-col :span="2" class="text-center">Action</el-col>
        </el-row>
        <div class="line"></div>
        <div class="user_agent" v-for="(user_agent, index) in configs.user_agents">
            <el-row :gutter="20">
                <el-col :span="3">
                    <el-input type="text" v-model.trim="user_agent.name"></el-input>
                </el-col>
                <el-col :span="19">
                    <el-input type="text" v-model.trim="user_agent.agent"></el-input>
                </el-col>
                <el-col :span="2" class="text-center">
                    <el-button type="danger" icon="delete" v-on:click="deleteUserAgent(index)"></el-button>
                </el-col>
            </el-row>
            <div class="line"></div>
        </div>

        <!--Create user agent dialog-->
        <el-dialog title="Performance" v-model="showCreateUserAgentForm">
            <el-form label-width="120px">
                <el-form-item label="ID">
                    <el-input type="text" v-model.trim="createUserAgentForm.userAgent.name"></el-input>
                </el-form-item>
                <el-form-item label="Proxy server">
                    <el-input type="text" v-model.trim="createUserAgentForm.userAgent.agent"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button v-on:click="showCreateUserAgentForm = false">Cancel</el-button>
                <el-button type="primary" v-on:click="createUserAgent">OK</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<style scoped>

</style>

<script>
export default {
    name: 'home-user-agents',

    props: {
        configs: {
            type: Object,
            required: true
        }
    },

    data() {
        return {
            showCreateUserAgentForm: false,
            createUserAgentForm: {
                userAgent: {
                    name: "",
                    agent: ""
                }
            }
        }
    },

    methods: {
        deleteUserAgent(id) {
            this.configs.user_agents = this.configs.user_agents.filter((user_agnet, index) => {
                return index != id
            })
        },

        createUserAgent() {
            // Clone object by simple way without _.clone(obj, true)
            const userAgent = JSON.parse(JSON.stringify(this.createUserAgentForm.userAgent))

            this.configs.user_agents.push(userAgent)
            this.resetCreateUserAgentForm()
            this.showCreateUserAgentForm = false
        },

        resetCreateUserAgentForm() {
            this.createUserAgentForm = {
                userAgent: {
                    name: "",
                    agent: ""
                }
            }
        }
    }
}
</script>
