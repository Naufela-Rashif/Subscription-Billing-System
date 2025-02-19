<template>
  <div class="lg:mb-20 lg:mt-0 sm:mt-20 mt-10 py-10 dark:bg-white">
    <Vue3Marquee :duration="10" direction="reverse" :pause-on-hover="true">
      <img
        v-for="(icon, index) in icons"
        :key="index"
        :src="icon"
        :alt="icon"
        class="lg:h-12 sm:h-10 h-5 lg:px-10 px-5 grayscale opacity-80 hover:grayscale-0 hover:opacity-100" />
    </Vue3Marquee>
  </div>
</template>

<script setup>
import { Vue3Marquee } from "vue3-marquee";
import youtube from "../assets/youtube.svg";
import wetv from "../assets/wetv.svg";
import viu from "../assets/viu.svg";
import spotify from "../assets/spotify.svg";
import netflix from "../assets/netflix.svg";
import disney from "../assets/disney.svg";
import capcut from "../assets/capcut.svg";
import applemusic from "../assets/applemusic.svg";

const icons = {
  youtube,
  wetv,
  viu,
  spotify,
  netflix,
  disney,
  capcut,
  applemusic,
};
</script>
