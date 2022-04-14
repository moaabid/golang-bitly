<script>
    import Item from './Item.svelte'
    import { onMount} from "svelte"

    let bitlies = []

    onMount(async ()=> {
        const res = await fetch("http://localhost:3000/getallbitlies")
        bitlies = await res.json()
    })

</script>

{#each bitlies as bitly }

<Item bitly = {bitly}/>

{/each}