<template>
    <div id="home-performances">
        <el-row :gutter="20" class="tab-heading">
            <el-col :span="3">
                <span class="tab-subject">Tasks</span>
            </el-col>
            <el-col :span="21" class="text-right">
                <el-button type="success" size="small" class="button-create" v-on:click="showCreatePerformanceForm = true">
                    <i class="el-icon-plus el-icon--left"></i>
                    Create
                </el-button>
            </el-col>
        </el-row>
        <div class="line"></div>

        <el-row :gutter="20">
            <el-col :span="2">ID</el-col>
            <el-col :span="3">Schedule</el-col>
            <el-col :span="5">Remark</el-col>
            <el-col :span="2" class="text-center">Enable</el-col>
            <el-col :span="2">Timeout</el-col>
            <el-col :span="2" class="text-center">Proxy</el-col>
            <el-col :span="3">Proxy Method</el-col>
            <el-col :span="3">Proxy Server</el-col>
            <el-col :span="2" class="text-center">Action</el-col>
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
                <el-col :span="5">
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
                <el-col :span="3">
                    <el-input type="text" v-model.trim="performance.proxy.server"></el-input>
                </el-col>
                <el-col :span="2" class="text-center">
                    <el-button type="danger" icon="delete" v-on:click="deletePerformance(performance.id)"></el-button>
                </el-col>
            </el-row>
            <div class="line"></div>
        </div>

        <!--Create performance dialog-->
        <el-dialog title="Performance" v-model="showCreatePerformanceForm">
            <el-form label-width="120px">
                <el-form-item label="ID">
                    <el-input type="number" v-model.number="createPerformanceForm.performance.id"></el-input>
                </el-form-item>
                <el-form-item label="Schedule">
                    <el-input type="type" v-model.trim="createPerformanceForm.performance.schedule"></el-input>
                </el-form-item>
                <el-form-item label="Remark">
                    <el-input type="type" v-model.trim="createPerformanceForm.performance.remark"></el-input>
                </el-form-item>
                <el-form-item label="Enable">
                    <el-checkbox v-model="createPerformanceForm.performance.enable"></el-checkbox>
                </el-form-item>
                <el-form-item label="Timeout">
                    <el-input type="number" v-model.number="createPerformanceForm.performance.timeout"></el-input>
                </el-form-item>
                <el-form-item label="Proxy enable">
                    <el-checkbox v-model="createPerformanceForm.performance.proxy.enable"></el-checkbox>
                </el-form-item>
                <el-form-item label="Proxy method">
                    <el-select v-model.trim="createPerformanceForm.performance.proxy.method">
                        <el-option label="Pool" value="pool"></el-option>
                        <el-option label="Custom" value="custom"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="Proxy server">
                    <el-input type="type" v-model.trim="createPerformanceForm.performance.proxy.server"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button v-on:click="showCreatePerformanceForm = false">Cancel</el-button>
                <el-button type="primary" v-on:click="showCreatePerformanceForm = false">OK</el-button>
            </div>
        </el-dialog>
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
    },

    data() {
        return {
            showCreatePerformanceForm: false,
            createPerformanceForm: {
                performance: {
                    id: 0,
                    schedule: "",
                    remark: "",
                    enable: false,
                    timeout: 3000,
                    proxy: {
                        enable: false,
                        method: "pool",
                        server: ""
                    }
                }
            }
        }
    },

    methods: {
        deletePerformance(id) {
            this.configs.performances = this.configs.performances.filter(performance => {
                return performance.id != id
            })
        }
    }
}
</script>
