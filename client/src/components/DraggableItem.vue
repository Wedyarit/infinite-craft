<script setup lang="ts">
import {useDrag} from 'vue3-dnd'
import {ItemTypes} from './ItemTypes'
import ItemCard from "@/components/ItemCard.vue";

const props = defineProps<{
  emoji: string
  title: string
}>()

const [, drag] = useDrag(() => ({
  type: ItemTypes.BOX,
  item: {title: props.title, emoji: props.emoji},
  collect: monitor => ({
    isDragging: monitor.isDragging(),
  }),
}))
</script>

<template>
  <div
      class="inline-block"
      :ref="drag"
      role="row"
      data-testid="box"
  >
    <ItemCard :id="'a'" :size="'small'" :title="title" :emoji="emoji"></ItemCard>
  </div>
</template>

<style scoped>

</style>