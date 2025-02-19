<template>
  <div
    class="lg:mt-0 mt-30 md:px-28 px-8 grid xl:grid-cols-3 md:grid-cols-2 grid-cols-2 md:gap-10 gap-3">
    <div
      class="border border-black rounded-2xl dark:border-white"
      v-for="product in products">
      <div class="flex flex-col items-center">
        <img
          :src="product.icon"
          alt=""
          class="w-full h-34 py-8 border-b border-black dark:bg-white rounded-2xl" />
        <h2
          class="capitalize mt-5 font-bold lg:text-3xl text-xl dark:text-white">
          {{ product.name }}
        </h2>
      </div>
      <div class="xl:px-10 px-2 py-5">
        <div
          class="flex xl:flex-row flex-col justify-between text-sm items-center dark:text-white">
          <p>Account Private U1</p>
          <p>{{ formatCurrency(product.price) }}</p>
        </div>
        <div
          class="mt-2 flex xl:flex-row flex-col justify-between text-sm items-center dark:text-white">
          <p>Account Sharing U4</p>
          <p>{{ formatCurrency(product.price + 30000) }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import youtube from "../assets/youtube.svg";
import wetv from "../assets/wetv.svg";
import viu from "../assets/viu.svg";
import spotify from "../assets/spotify.svg";
import netflix from "../assets/netflix.svg";
import disney from "../assets/disney.svg";
import capcut from "../assets/capcut.svg";
import applemusic from "../assets/applemusic.svg";
import { formatCurrency } from "../utils/currency";

const products = [
  { icon: youtube, name: "youtube", price: 20000 },
  { icon: wetv, name: "weTV", price: 20000 },
  { icon: viu, name: "viu", price: 20000 },
  { icon: spotify, name: "spotify", price: 20000 },
  { icon: netflix, name: "netflix", price: 20000 },
  { icon: disney, name: "disney", price: 20000 },
  { icon: capcut, name: "capcut", price: 20000 },
  { icon: applemusic, name: "apple music", price: 20000 },
];
</script>
