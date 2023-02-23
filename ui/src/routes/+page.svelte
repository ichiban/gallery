<script lang="ts">
    import type {Image} from '+page.ts'
    import lightGallery from 'lightgallery'
    import lgThumbnail from 'lightgallery/plugins/thumbnail'
    import lgZoom from 'lightgallery/plugins/zoom'
    import {onMount} from "svelte"

    export let data: Array<Image>;

    let gallery: HTMLDivElement

    onMount(() => {
        lightGallery(gallery, {
            plugins: [lgThumbnail, lgZoom],
            speed: 500,
        })
    })
</script>

<div class="gallery" bind:this={gallery}>
    {#each data.images || [] as image}
        <a href="{image.url}" data-lg-size="1600-2400">
            <img alt="" src="{image.thumb_url}" loading="lazy"/>
        </a>
    {/each}
</div>

<style>
    @import "lightgallery/css/lightgallery.css";
    @import "lightgallery/css/lg-zoom.css";
    @import "lightgallery/css/lg-thumbnail.css";

    :global(html) {
        background-color: #777870;
    }

    * {
        margin: 0;
        padding: 0;
    }

    .gallery {
        width: 100vw;
    }

    img {
        height: 100px;
        aspect-ratio: auto;
    }
</style>