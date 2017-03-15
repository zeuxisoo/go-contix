<template>
    <div id="home-user-agents">
        <h4>Agents</h4>
        <div class="line"></div>
        <el-row :gutter="20">
            <el-col :span="3">Name</el-col>
            <el-col :span="21">Agent</el-col>
        </el-row>
        <div class="line"></div>
        <div class="user_agent" v-for="user_agent in configs.user_agents">
            <el-row :gutter="20">
                <el-col :span="3">
                    <el-input type="text" v-model.trim="user_agent.name"></el-input>
                </el-col>
                <el-col :span="21">
                    <el-input type="text" v-model.trim="user_agent.agent"></el-input>
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
    }
}
</script>
