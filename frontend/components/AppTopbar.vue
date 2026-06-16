<script setup lang="ts">
import { Menu, Search } from '@lucide/vue'
</script>

<template>
  <header class="topbar">
    <div class="brand">
      <button class="icon-button" title="Menu">
        <Menu :size="20" />
      </button>
      <BrandCube />
      <span class="brand-title">BringIt</span>
    </div>

    <div class="search-pill">
      <div class="search-left">
        <Search :size="17" />
        <span class="search-text">Type a command or search</span>
      </div>
      <div class="kbd-group">
        <span class="kbd">Cmd</span>
        <span class="kbd">K</span>
      </div>
    </div>

    <div>
      <slot name="action" />
    </div>
  </header>
</template>
