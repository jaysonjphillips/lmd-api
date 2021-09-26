<template>
  <div class="mx-4">
    <table v-if="!loading && data && data.length" class="table table-striped">
        <tr>
            <th class="px-6" scope="col">ID</th>
            <th class="px-6" scope="col">Name</th>
            <th class="px-6" scope="col">Colorway</th>
            <th class="px-6" scope="col">Release Year</th>
            <th class="px-6" scope="col">Size</th>
            <th class="px-6" scope="col">Is Deadstock</th>
            <th class="px-6" scope="col">Is Sellable</th>
            <th class="px-6" scope="col">Is Packed</th>
            <th class="px-6" scope="col">Is Received</th>
            <th class="px-6" scope="col">Box Number</th>
        </tr>
        <tr v-for="sneaker of data" :key="sneaker.id">
          <td class="px-6">{{sneaker.id}}</td>
          <td class="px-6">{{sneaker.name}}</td>
          <td class="px-6">{{sneaker.colorway}}</td>
          <td class="px-6">{{sneaker.release_year}}</td>
          <td class="px-6">US {{sneaker.size}}</td>
          <td class="px-6">{{sneaker.is_deadstock}}</td>
          <td class="px-6">{{sneaker.is_sellable}}</td>
          <td class="px-6">{{sneaker.is_packed}}</td>
          <td class="px-6">{{sneaker.is_received}}</td>
          <td class="px-6">{{sneaker.box_number}}</td>

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
  name: 'SneakerList',
  props: {
    msg: String,
  },
  setup(): any {
    const data = ref(null);
    const loading = ref(true);
    const error = ref(null);

    function fetchData() {
      loading.value = true;

      return fetch('/api/v1/kicks')
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
