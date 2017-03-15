<template>
    <div id="home-mail">
        <el-row :gutter="20" class="tab-heading">
            <el-col :span="24">
                <span class="tab-subject">Mail</span>
            </el-col>
        </el-row>
        <div class="line"></div>

        <el-form ref="form" :model="configs.mail" label-width="80px">
            <el-form-item label="Sender">
                <el-input v-model.trim="configs.mail.sender"></el-input>
            </el-form-item>
            <el-form-item label="Recipient">
                <el-input v-model.trim="configs.mail.recipient"></el-input>
            </el-form-item>
            <el-form-item label="Subject">
                <el-input v-model.trim="configs.mail.subject"></el-input>
            </el-form-item>
        </el-form>
        <div class="line"></div>

        <el-row :gutter="20" class="tab-heading">
            <el-col :span="24">
                <span class="tab-subject">Mailgun</span>
            </el-col>
        </el-row>
        <div class="line"></div>

        <el-form ref="form" :model="configs.mail" label-width="80px">
            <el-form-item label="Domain">
                <el-input v-model.trim="configs.mail.mailgun.domain"></el-input>
            </el-form-item>
            <el-form-item label="Api Key">
                <el-input v-model.trim="configs.mail.mailgun.api_key"></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<style scoped>

</style>

<script>
export default {
    name: 'home-mail',

    props: {
        configs: {
            type: Object,
            required: true
        }
    }
}
</script>
