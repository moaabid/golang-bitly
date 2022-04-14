<script>
    import Card from "./Card.svelte"; 
    import { Modals, closeModal, openModal } from "svelte-modals"; 
    import Modal from "./Modal.svelte";
    export let bitly

    async function update(data){
        const json = {
            redirect: data.redirect,
            bitly: data.bitly,
            random: data.random,
            id: bitly.id
        }
        await fetch("http://localhost:3000/updatebitly", {
            method: "PATCH",
        headers: { "Content-Type":  "application/json"},
            body: JSON.stringify(json)
           
        }).then(res => {
            console.log(res)
        })
    }

    async function deleteBitly(id){
        if(confirm("Are you sure to delete (" + bitly.bitly+ ")")){
            await fetch(`http://localhost:3000/deletebitly/${id}`, {
           method: "DELETE" 
        }).then(res =>{
            console.log(res)
        })
        }
      
    }

function handleOpenModal(bitly){
    openModal(Modal, {
        title: "Update Bitly Link",
        send: update,
        bitly: bitly.bitly,
        redirect: bitly.redirect,
        random: bitly.random
    }) 
}
</script>

<Card>
    <p>Bitly: http://localhost:3000/r/{bitly.bitly}</p>
    <p>Redirect: {bitly.redirect}</p>
    <p>Clicked: {bitly.clicked}</p>
    <button class="update" on:click={handleOpenModal(bitly)}>Update</button>
    <button class="delete" on:click="{deleteBitly(bitly.id)}">Delete</button>
</Card>
<Modals>
    <div
      slot="backdrop"
      class="backdrop"
      on:click={closeModal}
    />
  </Modals>
<style>
    button{
        color: white;
        font-weight: bolder;
        border: none ;
        padding: .75rem;
        border-radius: 4px;
    }
    .update {
        background-color: rgb(76, 128, 249);
    }
    .delete{
         background-color: rgb(251, 89, 89);
    }
    .backdrop {
    position: fixed;
    top: 0;
    bottom: 0;
    right: 0;
    left: 0;
    background: rgba(0,0,0,0.50)
  }
</style>