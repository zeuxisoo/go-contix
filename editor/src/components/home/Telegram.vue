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

        <el-row :gutter="20" class="tab-heading">
            <el-col :span="3">
                <span class="tab-subject">Chat Ids</span>
            </el-col>
            <el-col :span="21" class="text-right">
                <el-button type="success" size="small" class="button-create" v-on:click="showCreateChatIdForm = true">
                    <i class="el-icon-plus el-icon--left"></i>
                    Create
                </el-button>
            </el-col>
        </el-row>
        <div class="line"></div>

        <el-row :gutter="20">
            <el-col :span="3">Name</el-col>
            <el-col :span="19">Code</el-col>
            <el-col :span="2" class="text-center">Action</el-col>
        </el-row>
        <div class="line"></div>
        <div class="chat_id" v-for="(chat_id, index) in configs.telegram.chat_ids">
            <el-row :gutter="20">
                <el-col :span="3">
                    <el-input type="text" v-model.trim="chat_id.name"></el-input>
                </el-col>
                <el-col :span="19">
                    <el-input type="number" v-model.number="chat_id.code"></el-input>
                </el-col>
                <el-col :span="2" class="text-center">
                    <el-button type="danger" icon="delete" v-on:click="deleteChatId(index)"></el-button>
                </el-col>
            </el-row>
            <div class="line"></div>
        </div>

        <!--Create chat id dialog-->
        <el-dialog title="Performance" v-model="showCreateChatIdForm">
            <el-form label-width="120px">
                <el-form-item label="Name">
                    <el-input type="text" v-model.trim="createChatIdForm.chatId.name"></el-input>
                </el-form-item>
                <el-form-item label="Code">
                    <el-input type="number" v-model.number="createChatIdForm.chatId.code"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button v-on:click="showCreateChatIdForm = false">Cancel</el-button>
                <el-button type="primary" v-on:click="createChatId">OK</el-button>
            </div>
        </el-dialog>
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
    },

    data() {
        return {
            showCreateChatIdForm: false,
            createChatIdForm: {
                chatId: {
                    name: "",
                    code: ""
                }
            }
        }
    },

    methods: {
        deleteChatId(id) {
            this.configs.telegram.chat_ids = this.configs.telegram.chat_ids.filter((chat_id, index) => {
                return index != id
            })
        },

        createChatId() {
            // Clone object by simple way without _.clone(obj, true)
            const chatId = JSON.parse(JSON.stringify(this.createChatIdForm.chatId))

            this.configs.telegram.chat_ids.push(chatId)
            this.resetCreateChatIdForm()
            this.showCreateChatIdForm = false
        },

        resetCreateChatIdForm() {
            this.createChatIdForm = {
                chatId: {
                    name: "",
                    code: ""
                }
            }
        }
    }
}
</script>
