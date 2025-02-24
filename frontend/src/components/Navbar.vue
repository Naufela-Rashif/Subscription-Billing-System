<template>
  <nav
    :class="
      y > 30
        ? 'bg-black text-white dark:bg-black py-5 shadow-lg shadow-black/20 dark:shadow-white/10'
        : 'py-10'
    "
    class="sm:px-16 xl:px-28 px-8 flex w-full flex-row justify-between items-center z-[99] fixed transition-all duration-300">
    <h2 class="lg:text-3xl text-xl font-black dark:text-white">
      akunSharing.id
    </h2>
    <div class="hidden sm:visible sm:flex items-center md:gap-x-10 gap-x-5">
      <a
        v-for="(page, index) in pages"
        :key="index"
        href=""
        class="lg:text-lg text-sm font-medium dark:text-white dark:hover:underline dark:hover:underline-offset-4"
        ><p>{{ page.name }}</p></a
      >
      <div class="h-10 flex items-center gap-x-2">
        <a
          :class="y > 30 ? 'border-white' : 'border-black'"
          href=""
          class="h-full flex items-center border rounded-xl hover:bg-lime-300 dark:text-white dark:border-white dark:hover:bg-lime-300/50"
          ><p class="lg:text-lg text-sm md:px-6 px-3">Hubungi</p></a
        >
        <img
          :src="isDark ? moon : sun"
          @click="toogleDark()"
          alt=""
          :class="
            y > 30
              ? 'hover:bg-white/20 invert brightness-0'
              : 'border-black hover:bg-black/10'
          "
          class="h-full p-2 rounded-xl border cursor-pointer dark:border-white dark:hover:bg-white/20" />
      </div>
    </div>
    <div class="relative sm:hidden">
      <!-- Hamburger Button -->
      <button @click="isOpen = !isOpen" class="flex flex-col space-y-1 p-2">
        <span
          class="w-8 h-1 bg-black transition dark:bg-white rounded"
          :class="{ 'rotate-45 translate-y-2': isOpen }"></span>
        <span
          class="w-8 h-1 bg-black transition dark:bg-white rounded"
          :class="{ 'opacity-0': isOpen }"></span>
        <span
          class="w-8 h-1 bg-black transition dark:bg-white rounded"
          :class="{ '-rotate-45 -translate-y-2': isOpen }"></span>
      </button>

      <!-- Menu -->
      <div v-if="isOpen" class="absolute left-0 top-12 w-40 bg-white shadow-md">
        <a href="#" class="block p-2 hover:bg-gray-100">Home</a>
        <a href="#" class="block p-2 hover:bg-gray-100">About</a>
        <a href="#" class="block p-2 hover:bg-gray-100">Services</a>
        <a href="#" class="block p-2 hover:bg-gray-100">Contact</a>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from "vue";
import { useDark, useToggle, useWindowScroll } from "@vueuse/core";
const { y } = useWindowScroll();
import moon from "../assets/moon.svg";
import sun from "../assets/sun.svg";

let isOpen = ref(false);
const isDark = useDark();
const toogleDark = useToggle(isDark);

let pages = [
  {
    name: "Home",
    link: "",
  },
  {
    name: "Product",
    link: "",
  },
  {
    name: "Cara Pesan",
    link: "",
  },
];
</script>
