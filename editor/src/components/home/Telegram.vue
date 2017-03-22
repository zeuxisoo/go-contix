<template>
    <div id="home-mail">
        <el-row :gutter="20" class="tab-heading">
            <el-col :span="24">
                <span class="tab-subject">Telegram</span>
            </el-col>
        </el-row>
        <div class="line"></div>

        <el-form ref="form" :model="configs.telegram" label-width="80px">
            <el-form-item label="Enable">
                <el-checkbox v-model="configs.telegram.enable"></el-checkbox>
            </el-form-item>
            <el-form-item label="Token">
                <el-input v-model.trim="configs.telegram.token"></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<style scoped>

</style>

<script>
export default {
    name: 'home-telegram',

    props: {
        configs: {
            type: Object,
            required: true
        }
    }
}
</script>
