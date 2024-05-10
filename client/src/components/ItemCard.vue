<script setup lang="ts">
import {useDrop} from "vue3-dnd";
import {ItemTypes} from "@/components/ItemTypes";
import {type DragItem} from "@/components/interfaces";
import {useBoxesStore} from "@/stores/useBoxesStore";
import axios from "axios";
import {useResourcesStore} from "@/stores/useResourcesStore";
import {storeToRefs} from "pinia";
import {twMerge} from "tailwind-merge";
import newResourceSound from '@/assets/sounds/new_resource_sound.mp3';
import {nextTick} from "vue";
const newResourceAudio = new Audio(newResourceSound);

const playNewResourceSound = () => {
  newResourceAudio.play();
};

const props = defineProps<{
  title: string;
  emoji: string;
  id: string;
  size: 'small' | 'large';
}>()

const store = useBoxesStore()
const {removeBox, addBox} = store
const {resources} = storeToRefs(useResourcesStore())
const {addResource} = useResourcesStore()

const [, drop] = useDrop(() => ({
  accept: ItemTypes.BOX,
  async drop(item: DragItem) {
    if (item.id !== props.id) {
      const droppedId = item?.id;
      const secondTitle = store.boxes[droppedId]?.title ?? item?.title
      if (droppedId) {
        removeBox(droppedId);
      }
      store.boxes[props.id].loading = true
      try {
        const response = await axios.get(`http://127.0.0.1:3000/api/recipe?first_ingredient=${store.boxes[props.id].title}&second_ingredient=${secondTitle}`)
        const resultAnswer = response.data.result || store.boxes[props.id].title
        const resultEmoji = response.data.emoji || store.boxes[props.id].emoji

        const newBoxId = addBox({
          title: resultAnswer,
          emoji: resultEmoji,
          left: store.boxes[props.id].left,
          top: store.boxes[props.id].top
        })
        if (!resources.value.find((resource: { title: string; }) => resource.title === resultAnswer)) {
          addResource({
            title: resultAnswer,
            emoji: resultEmoji
          })
          await nextTick(() => {
            playNewResourceSound();
            playAnimation(newBoxId);
          });

        }
      } catch (error) {
        console.error('Error fetching recipe:', error)
      }
      removeBox(props.id);
    }
  },
}))

const playAnimation = (boxId: string) => {
  const boxElement = document.getElementById(`${boxId}`);

  if (boxElement) {
    const particle = boxElement.querySelector('.particle');
    if (particle) {
      particle.classList.toggle('particles');
      particle.classList.toggle('animated');
      setTimeout(() => {
        particle.classList.toggle('animated');
        particle.classList.toggle('particles');
      }, 1660);
    }
  }
}

</script>

<template>
  <div class='buttonWrap'>
    <div :ref="drop"
         :class="twMerge(props.size === 'large' ? 'text-2xl space-x-2.5 py-2.5 px-4' : 'space-x-1.5 px-3 py-1','border-gray-200 bg-white shadow hover:bg-gray-100 cursor-pointer transition inline-block font-medium border rounded-lg theButton particleButton animated')">
          <span>
          {{ emoji }}
        </span>
      <span>
        {{ title }}
      </span>
    </div>
    <div class='theButton particleButton particle'></div>
  </div>
</template>

<style scoped>
</style>