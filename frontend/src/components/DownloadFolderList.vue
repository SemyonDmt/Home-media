<script setup lang="ts">

defineProps<{ folders: string[] }>();

const emit = defineEmits<{
  save: [value: string],
  close: []
}>()

const onDownloadToFolder = (name: string) => emit('save', name);
const onClose = () => emit('close')
</script>

<template>
  <div class="modal fade" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered">
      <div class="modal-content">
        <div class="modal-body">
          <div class="container">
            <div class="row gy-3">
              <h5 class="col modal-title text-center" id="exampleModalLabel">Выберите папку для загрузки</h5>
              <button type="button" class="col-6 offset-3 btn btn-primary" v-for="folder in folders" :key="folder"
                      @click="onDownloadToFolder(folder)">{{ folder }}
              </button>
              <button type="button" class="col-6 offset-3 btn btn-secondary" @click="onClose">Отменить</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">

</style>