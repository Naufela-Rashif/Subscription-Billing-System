<script setup lang="ts">
import Logos from "./Logos.vue";
</script>

<template>
  <div class="lg:h-screen lg:pt-10 pt-30 flex flex-col justify-between">
    <div
      class="sm:px-16 xl:px-28 px-8 flex items-center h-full lg:gap-x-20 gap-x-5">
      <div class="w-[50%]">
        <h2
          class="xl:text-6xl xl:leading-20 lg:text-5xl md:text-3xl lg:leading-16 sm:leading-10 text-balance text-2xl font-black dark:text-lime-300">
          Ayo percayakan akun premium anda pada kami
        </h2>
        <p
          class="mt-10 lg:text-xl text-sm text-balance text-[#E3651D] dark:text-white">
          akunSharing.id menawarkan beberapa akun premium seperti akun sharing,
          hingga private, tersedia untuk multi user seperti tv, hp, laptop,
          dengan jaminan garansi dan layanan 24/7 jam. Ayo pesan sekarang!
        </p>
        <a href=""
          ><p
            class="mt-8 bg-lime-300 border border-black text-black font-bold xl:text-md sm:px-14 sm:py-5 py-3 px-8 inline-block rounded-xl dark:text-black">
            Pesan Sekarang!
          </p></a
        >
      </div>
      <h1 class="dark:text-white">Image</h1>
    </div>
    <Logos />
  </div>
</template>
