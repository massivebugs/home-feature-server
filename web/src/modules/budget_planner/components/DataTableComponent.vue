<template>
  <div>
    <DataTable :data="data" class="display"> </DataTable>
  </div>
</template>
<script setup lang="ts">
import DataTable from 'datatables.net-vue3'
import DataTablesCore from 'datatables.net'
import 'datatables.net-select'
import 'datatables.net-responsive'
import 'datatables.net-buttons'

DataTable.use(DataTablesCore)

const columns = [{}]

const data = [
  [1, 2],
  [3, 4],
]
</script>
<style scoped lang="scss">
@use 'xp.css/dist/XP.css';
@use '@/assets/xp.custom';
</style>
