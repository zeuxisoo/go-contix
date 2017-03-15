<template>
    <div id="home-performances">
        <h4>Tasks</h4>
        <div class="line"></div>

        <el-row :gutter="20">
            <el-col :span="2">ID</el-col>
            <el-col :span="3">Schedule</el-col>
            <el-col :span="6">Remark</el-col>
            <el-col :span="2" class="text-center">Enable</el-col>
            <el-col :span="2">Timeout</el-col>
            <el-col :span="2" class="text-center">Proxy</el-col>
            <el-col :span="3">Proxy Method</el-col>
            <el-col :span="4">Proxy Server</el-col>
        </el-row>
        <div class="line"></div>

        <div class="performance" v-for="performance in configs.performances">
            <el-row :gutter="20">
                <el-col :span="2">
                    <el-input type="number" v-model.number="performance.id"></el-input>
                </el-col>
                <el-col :span="3">
                    <el-input type="text" v-model.trim="performance.schedule"></el-input>
                </el-col>
                <el-col :span="6">
                    <el-input type="text" v-model.trim="performance.remark"></el-input>
                </el-col>
                <el-col :span="2" class="checkbox-center">
                    <el-checkbox v-model="performance.enable"></el-checkbox>
                </el-col>
                <el-col :span="2">
                    <el-input type="number" v-model.number="performance.timeout"></el-input>
                </el-col>
                <el-col :span="2" class="checkbox-center">
                    <el-checkbox v-model="performance.proxy.enable"></el-checkbox>
                </el-col>
                <el-col :span="3">
                    <el-input type="text" v-model.trim="performance.proxy.method"></el-input>
                </el-col>
                <el-col :span="4">
                    <el-input type="text" v-model.trim="performance.proxy.server"></el-input>
                </el-col>
            </el-row>
            <div class="line"></div>
        </div>
    </div>
</template>

<style scoped>
.checkbox-center {
    margin-top: 8px;
    text-align: center;
}
</style>

<script>
export default {
    name: 'home-performances',

    props: {
        configs: {
            type: Object,
            required: true
        }
    }
}
</script>
