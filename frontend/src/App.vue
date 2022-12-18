<template>
  <div class="toast-container position-fixed top-0 start-50 translate-middle-x pt-3">
    <div class="toast align-items-center text-bg-primary border-0" id="toastForSuccess" role="status"
         aria-live="assertive" aria-atomic="true">
      <div class="d-flex">
        <div class="toast-body">
          Успешно поставлен в загрузку.
        </div>
        <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast"
                aria-label="Close"></button>
      </div>
    </div>
    <div class="toast align-items-center text-bg-danger border-0" id="toastForError" role="status" aria-live="assertive"
         aria-atomic="true">
      <div class="d-flex">
        <div class="toast-body">
          Не удалось поставить в загрузку. Ошибка: {{ errorMessage }}
        </div>
        <button type="button" class="btn-close btn-close-white me-2 m-auto" data-bs-dismiss="toast"
                aria-label="Close"></button>
      </div>
    </div>
  </div>

  <div class="container pt-5">
    <header class="row pt-3">
      <div class="col text-center">Найти фильм</div>
    </header>
    <form v-on:submit.prevent="search" class="row pt-4">
      <div class="col-12">
        <input v-model="movieName" type="text" id="website-input" placeholder="Введите название фильма"
               class="form-control">
      </div>
      <div class="col-12 gy-1">
        <input v-model="movieYear" type="number" id="website-input" placeholder="Введите год фильма"
               class="form-control">
      </div>
      <div class="col-12 text-center gy-3">
        <button class="btn btn-primary" :disabled="isDownload">Найти</button>
      </div>
    </form>
    <section class="row pt-3">
      <div class="col text-center">
        <div class="d-inline-block form-check">
          <input v-model="more10gb" class="form-check-input" type="checkbox" id="check-more-10GB">
          <label class="form-check-label" for="check-more-10GB">&lt; 10 GB</label>
        </div>
        <div class="d-inline-block ms-3 form-check">
          <input v-model="between1and10gb" class="form-check-input" type="checkbox" id="check-more-1GB">
          <label class="form-check-label" for="check-more-1GB">1-10 GB</label>
        </div>
        <div class="d-inline-block ms-3 form-check">
          <input v-model="less1gb" class="form-check-input" type="checkbox" id="check-less-1GB">
          <label class="form-check-label" for="check-less-1GB">&gt; 1 GB</label>
        </div>
      </div>
    </section>
    <section class="row pt-3">
      <div class="col text-center">
        <div v-if="isDownload" class="spinner-border"></div>
        <div v-if="getItems.length === 0 && !isNew && !isDownload" class="text-center">Ни чего не найдено</div>
      </div>
    </section>
    <section class="row">
      <div v-if="getItems.length > 0" class="col">
        <table class="table">
          <thead class="text-center">
          <tr>
            <th scope="col"></th>
            <th scope="col">
              <i class="bi bi-person-fill-down"></i>
            </th>
            <th scope="col">
              <i class="bi bi-hdd-fill"></i>
            </th>
            <th scope="col"></th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="item in getItems" :key="item">
            <td>
              <i class="btn btn-primary bi bi-download" @click="download(item)"></i>
            </td>
            <td>{{ item.seeders }}</td>
            <td>{{ toShortSize(item.size) }}</td>
            <td>
              <div style="margin-left: 15px">
                <a :href="item.details" target="_blank">{{ item.title }}</a>
              </div>
            </td>
          </tr>
          </tbody>
        </table>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import axios, {AxiosError, isAxiosError} from 'axios';
import {Toast} from "bootstrap";
import {computed, onMounted, ref} from 'vue'
import {Torrent} from "@/models/torrent";
import {gb, tenGb, toShortSize} from "@/helper/size";
import {DownloadResult} from "@/models/DownloadResult";

const movieName = ref<string>('');
const movieYear = ref<string>('');
const isDownload = ref<boolean>(false);
const isNew = ref<boolean>(true);
const more10gb = ref<boolean>(true);
const between1and10gb = ref<boolean>(true);
const less1gb = ref<boolean>(true);
const items = ref<Torrent[]>([])
let toastSuccess: Toast | null = null;
let toastError: Toast | null = null;
const errorMessage = ref<string | unknown>('');

onMounted(() => {
  toastSuccess = new Toast('#toastForSuccess');
  toastError = new Toast('#toastForError', {autohide: false});
})

const getItems = computed(() => {
  return items.value.filter(({size}) => {
    if (less1gb.value && size < gb) {
      return true
    }
    if (between1and10gb.value && size >= gb && size < tenGb) {
      return true
    }
    return more10gb.value && size >= tenGb;

  })
})

async function search(): Promise<void> {
  isDownload.value = true
  items.value = []
  isNew.value = false
  try {
    const {data} = await axios.post<Torrent[]>("/api/items", {search: `${movieName.value} ${movieYear.value}`})
    items.value = data;
  } catch (error) {
    const err = isAxiosError(error) ? (error as AxiosError).message : error;
    showToastError(err);
  } finally {
    isDownload.value = false
  }
}

async function download(torrent: Torrent): Promise<void> {
  try {
    const request = {trackerId: torrent.trackerId, downloadLink: torrent.downloadLink};
    const {data} = await axios.post<DownloadResult>("/api/download", request);
    if (data.result) {
      showToastSuccess()
    } else {
      showToastError(data.errorMessage)
    }
  } catch (error) {
    const err = isAxiosError(error) ? (error as AxiosError).message : error;
    showToastError(err);
  }
}

function showToastError(message: string | unknown): void {
  errorMessage.value = message;
  toastError?.show();
}

function showToastSuccess(): void {
  toastSuccess?.show();
}
</script>

<style lang="scss">
.container {
  header {
    font-size: 1.75em;
    font-weight: 500;
  }

  .table {
    font-size: 14px;

    thead {
      font-size: 22px;
      border-top: solid 1px rgba(222, 226, 230, 0.75);
    }

    a {
      color: #428bca;
    }
  }
}
</style>
