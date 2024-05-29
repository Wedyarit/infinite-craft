<script lang="ts" setup>
import {useDrop, type XYCoord} from 'vue3-dnd'
import {ItemTypes} from './ItemTypes'
import DraggableBox from './DraggableBox.vue'
import type {DragItem} from './interfaces'
import {computed, ref} from 'vue'
import ItemCard from "@/components/ItemCard.vue";
import AvailableResources from "@/components/AvailableResources.vue";
import {useBoxesStore} from "@/stores/useBoxesStore";
import LeaderBoard from "@/components/LeaderBoard.vue";

const store = useBoxesStore()
const {boxes} = store

const moveBox = (id: string | null, left: number, top: number, title: string, emoji: string) => {
  if (id) {
    Object.assign(boxes[id], {left, top})
  } else {
    const key = Math.random().toString(36).substring(7)
    boxes[key] = {top, left, title, emoji}
  }
}

const containerElement = ref<HTMLElement | null>(null)

const [, drop] = useDrop(() => ({
  accept: ItemTypes.BOX,
  drop(item: DragItem, monitor) {
    if (item.id && item.left !== null && item.top !== null) {
      const delta = monitor.getDifferenceFromInitialOffset() as XYCoord
      if (delta && delta.x && delta.y) {
        const left = Math.round((item.left) + delta.x)
        const top = Math.round((item.top) + delta.y)
        moveBox(item.id, left, top, item.title, item.emoji)
      }
    } else {
      const delta = monitor.getClientOffset() as XYCoord
      const containerCoords = containerElement?.value?.getBoundingClientRect()
      if (containerCoords == null) return;
      if (delta && delta.x && delta.y) {
        const left = Math.round(delta.x - containerCoords.left - 40)
        const top = Math.round(delta.y - containerCoords.top - 15)
        moveBox(null, left, top, item.title, item.emoji)
      }
    }
  },
}))

const selectedTab = ref('resources') // State to track selected tab

const activeComponent = computed(() => {
  return selectedTab.value === 'resources' ? AvailableResources : LeaderBoard
})
</script>

<template>
  <div ref="containerElement">
    <main class="flex gap-x-3">
      <div class="w-3/4">
        <div :ref="drop" class="container">
          <DraggableBox
              v-for="(value, key) in boxes"
              :id="key"
              :key="key"
              :left="value.left"
              :top="value.top"
              :loading="value.loading"
          >
            <ItemCard size="large" :id="key.toString()" :title="value.title" :emoji="value.emoji"/>
          </DraggableBox>
        </div>
      </div>
      <div class="w-1/4 bg-white shadow px-4 py-3 border-gray-200 border rounded-lg overflow-y-scroll max-h-[80vh]">
        <div class="flex gap-x-2 mb-4">
          <button @click="selectedTab = 'resources'" :class="{ 'font-bold': selectedTab === 'resources' }"
                  class="tab-button">Ресурсы
          </button>
          <button @click="selectedTab = 'leaders'" :class="{ 'font-bold': selectedTab === 'leaders' }"
                  class="tab-button">Лидеры
          </button>
        </div>
        <h2 class="font-semibold">{{ selectedTab === 'resources' ? 'Ресурсы' : 'Лидеры' }}</h2>
        <component :is="activeComponent"></component>
      </div>
    </main>
  </div>
</template>

<style scoped>
.container {
  position: relative;
  width: 100%;
  height: 90vh;
}

.tab-button {
  flex: 1;
  padding: 12px;
  font-size: 1.1rem;
  text-align: center;
  background-color: #f3f3f3;
  border: none;
  cursor: pointer;
  border-radius: 12px;
  transition: background-color 0.3s ease;
}

.tab-button:hover {
  background-color: #e3e3e3;
}

.tab-button.font-bold {
  font-weight: bold;
  background-color: #d3d3d3;
}
</style>
