<template>
  <v-app>
    <div>pending: {{ pending }}</div>
    <div>error: {{ error }}</div>
    <v-btn color="primary" @click="refresh">リフレッシュ</v-btn>
    <v-table density="compact">
      <thead>
      <tr>
        <th class="text-left">id</th>
        <th class="text-left">displayName</th>
        <th class="text-left">fullName</th>
        <th class="text-left">firstName</th>
        <th class="text-left">lastName</th>
        <th class="text-left">birthday</th>
        <th class="text-left">age</th>
        <th class="text-left">registeredAt</th>
        <th class="text-left">latestUpdatedAt</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="user in users" :key="user.id">
        <td>{{ user.id }}</td>
        <td>{{ user.displayName }}</td>
        <td>{{ user.fullName }}</td>
        <td>{{ user.firstName }}</td>
        <td>{{ user.lastName }}</td>
        <td>{{ user.birthday }}</td>
        <td>{{ user.age }}</td>
        <td>{{ user.registeredAt }}</td>
        <td>{{ user.latestUpdatedAt }}</td>
      </tr>
      </tbody>
    </v-table>
  </v-app>
</template>
<script setup lang="ts">
import {useFetch} from "#imports";

const { data: users, refresh, error, pending } = await useFetch<User[]>('http://localhost:5050/v1/user')
</script>
