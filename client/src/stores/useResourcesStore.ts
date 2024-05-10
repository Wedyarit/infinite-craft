import {defineStore} from 'pinia'
import {useLocalStorage} from "@vueuse/core";

export interface ResourceStoreEntry {
    title: string
    emoji: string
}

export const useResourcesStore = defineStore('resources', () => {
    const resources =
        useLocalStorage<ResourceStoreEntry[]>('opencraft/resources', [
            {title: 'Огонь', emoji: '🔥'},
            {title: 'Вода', emoji: '💧'},
            {title: 'Земля', emoji: '🌍'},
            {title: 'Воздух', emoji: '💨'},
        ]);

    function addResource(box: ResourceStoreEntry) {
        resources.value.push(box)
    }

    return {resources, addResource}
})
