<template>
    <div id="home-user-agents">
        <el-row :gutter="20" class="tab-heading">
            <el-col :span="3">
                <span class="tab-subject">Agents</span>
            </el-col>
            <el-col :span="21" class="text-right">
                <el-button type="success" size="small" class="button-create" v-on:click="">
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

    methods: {
        deleteUserAgent(id) {
            this.configs.user_agents = this.configs.user_agents.filter((user_agnet, index) => {
                return index != id
            })
        }
    }
}
</script>
