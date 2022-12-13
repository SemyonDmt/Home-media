<style>
html {
  font-size: 100%;
}
</style>

<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1 class="text-center">Найти фильм</h1>
        <br>
        <form v-on:submit.prevent="search">
          <div class="form-group">
            <input v-model="movieName" type="text" id="website-input" placeholder="Введите название фильма"
                   class="form-control">
            <input v-model="movieYear" type="number" id="website-input" placeholder="Введите год фильма"
                   class="form-control">
          </div>
          <div class="d-flex justify-content-center form-group">
            <button class="btn btn-primary" :disabled="isDownload">Найти</button>
          </div>
          <div class="row d-flex flex-row justify-content-center">
            <div class="form-check">
              <input v-model="more10gb" class="form-check-input" type="checkbox" value="" id="check-more-10GB">
              <label class="form-check-label" for="check-more-10GB">&lt; 10 GB</label>
            </div>
            <div style="margin-left: 1em" class="form-check">
              <input v-model="between1and10gb" class="form-check-input" type="checkbox" value="" id="check-more-1GB">
              <label class="form-check-label" for="check-more-1GB">1-10 GB</label>
            </div>
            <div style="margin-left: 1em" class="form-check">
              <input v-model="less1gb" class="form-check-input" type="checkbox" value="" id="check-less-1GB">
              <label class="form-check-label" for="check-less-1GB">&gt; 1 GB</label>
            </div>
          </div>
        </form>
        <div v-if="isDownload" class="d-flex justify-content-center">
          <div class="spinner-border text-secondary" role="status">
            <span class="sr-only">Loading...</span>
          </div>
        </div>
      </div>
      <div class="col-md-12">
        <div>
          <div v-if="getItems.length === 0 && !isNew && !isDownload" class="text-center">Ни чего не найдено</div>
          <table v-if="getItems.length > 0" class="table table-sm">
            <tr>
              <th scope="col"></th>
              <th scope="col" class="text-center">
                <!--               ToDo заменить на иконку https://icons.getbootstrap.com/icons/person-fill-down/-->
                <span class="glyphicon glyphicon-user" aria-hidden="true"></span>
              </th>
              <th scope="col" style="width: 5em;" class="text-center">
                <!--              <th scope="col" class="text-center">-->
                <span class="glyphicon glyphicon-hdd" aria-hidden="true"></span>
              </th>
              <th scope="col"></th>
            </tr>
            <tr v-for="item in getItems" :key="item">
              <td>
                <button type="button" class="btn btn-primary" v-on:click="download(item)">
                  <span class="glyphicon glyphicon-download-alt" aria-hidden="true"></span>
                </button>
              </td>
              <td class="text-center">{{ item.seeders }}</td>
              <td class="text-right">{{ toShortSize(item.size) }}</td>
              <td>
                <div style="margin-left: 15px">
                  <a :href="item.details" target="_blank">{{ item.title }}</a>
                </div>

              </td>
            </tr>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import {kb, mb, gb, tb, tenGb} from "@/helper/size";

export default {
  name: 'App',

  data() {
    return {
      movieName: '',
      movieYear: '',
      items: [],
      isDownload: false,
      isNew: true,
      more10gb: true,
      between1and10gb: true,
      less1gb: true,
    }
  },

  computed: {
    getItems() {
      return this.items.filter(({size}) => {
        if (this.less1gb && size < gb) {
          return true
        }
        if (this.between1and10gb && size >= gb && size < tenGb) {
          return true
        }
        if (this.more10gb && size >= tenGb) {
          return true
        }
        return false;
      })
    }
  },

  methods: {
    search() {
      this.isDownload = true
      this.items = []
      this.isNew = false
      axios.post("/api/items", {search: `${this.movieName} ${this.movieYear}`})
          .then((response) => {
            this.items = response.data;
            this.isDownload = false
          })
          .catch((error) => {
            window.alert(`The API returned an error: ${error}`);
            this.isDownload = false
          })
    },
    download(item) {
      axios.post("/api/download", {trackerId : item.trackerId, downloadLink: item.downloadLink})
          .then((response) => {
            console.log(response)
          })
          .catch((error) => {
            window.alert(`The API returned an error: ${error}`);
          })
    },
    toShortSize(size) {
      if (size < kb) {
        return `${size} B`
      }
      if (size < mb) {
        const s = size / kb < 100 ? (size / kb).toFixed(1) : (size / kb).toFixed(0)
        return `${s} KB`
      }
      if (size < gb) {
        const s = size / mb < 100 ? (size / mb).toFixed(1) : (size / mb).toFixed(0)
        return `${s} MB`
      }
      if (size < tb) {
        const s = size / gb < 100 ? (size / gb).toFixed(1) : (size / gb).toFixed(0)
        return `${s} GB`
      }
      const s = size / tb < 100 ? (size / tb).toFixed(1) : (size / tb).toFixed(0)
      return `${s} TB`
    }
  }
}
</script>