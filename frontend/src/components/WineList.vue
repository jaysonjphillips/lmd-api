<template>
  <div class="mx-4">
    <table v-if="!loading && data && data.length" class="table table-striped">
        <tr>
            <th class="px-6" scope="col">ID</th>
            <th class="px-6" scope="col">Name</th>
            <th class="px-6" scope="col">Vineyard</th>
            <th class="px-6" scope="col">Location</th>
            <th class="px-6" scope="col">Year</th>
            <th class="px-6" scope="col">Is Reserve</th>
            <th class="px-6" scope="col">Is Stored</th>
            <th class="px-6" scope="col">Is Packed</th>
            <th class="px-6" scope="col">Is Received</th>
            <th class="px-6" scope="col">Box Number</th>
        </tr>
        <tr v-for="wine of data" :key="wine.id">
          <td class="px-6">{{wine.id}}</td>
          <td class="px-6">{{wine.name}}</td>
          <td class="px-6">{{wine.vineyard}}</td>
          <td class="px-6">{{wine.location}}</td>
          <td class="px-6">{{wine.year}}</td>
          <td class="px-6">{{wine.is_reserve}}</td>
          <td class="px-6">{{wine.is_stored}}</td>
          <td class="px-6">{{wine.is_packed}}</td>
          <td class="px-6">{{wine.is_received}}</td>
          <td class="px-6">{{wine.box_number}}</td>

        </tr>
    </table>
    <p v-if="loading">
      Still loading..
    </p>
    <p v-if="error">

    </p>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue';

export default {
  name: 'WineList',
  props: {
    msg: String,
  },
  setup(): any {
    const data = ref(null);
    const loading = ref(true);
    const error = ref(null);

    function fetchData() {
      loading.value = true;

      return fetch('/api/v1/wines')
        .then((res) => {
          if (!res.ok) {
            // do something
          }

          return res.json();
        })
        .then((json) => {
          console.log(json.data);
          data.value = json.data;
          loading.value = false;
        });
      // Will be implemented next
    }

    onMounted(() => {
      fetchData();
    });

    return {
      data,
      loading,
      error,
    };
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  th, td {
    padding: 5px 10px;
  }
</style>
