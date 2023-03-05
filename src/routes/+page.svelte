<script>
    import { onMount } from "svelte";
    import { _data, _read_data, _newTask, _updateStores, _deleteTask, _updateStatus} from "./+page"

    let _taskDesc

    onMount(async () => {
        fetch("https://todolist-zillaz.koyeb.app/")
        .then(response => response.json())
        .then(_read_data => {
            console.log(_read_data)
            _data.set(_read_data)
        }).catch(error => {
            console.log(error)
            return []
        })
    })
</script>

<main>
    <h1> Todo App </h1>
    <input type="text" bind:value={_taskDesc}/>
    <button on:click= { _newTask(_taskDesc) }> Add Task </button>
    <div>
        {#each $_data as d}
        {#if d.Completed }
        <del> { d.Description } </del>
        <button on:click= {_updateStatus(d.Description)}> Uncheck </button>
        <button on:click= {_deleteTask(d.Description)}> Delete Task </button>
        {:else}
        <p> { d.Description } </p>
        <button on:click= {_updateStatus(d.Description)}> Check </button> <br>
        {/if}
        {/each}
    </div>
</main>