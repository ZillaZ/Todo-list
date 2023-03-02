<script>
    import { onMount } from "svelte";
    import { _data, _read_data, _newTask} from "./+page"

    let taskDesc;

    onMount(async () => {
        fetch("http://127.0.0.1:5050")
        .then(response => response.json())
        .then(_read_data => {
            console.log(_read_data)
            _data.set(_read_data)
        }).catch(error => {
            console.log(error)
            return []
        })
    })

    function updateStatus(desc) {
        fetch("http://127.0.0.1:5050/" + desc)
        updateStores()
    }

    function updateStores() {
        onMount(async () => {
            fetch("http://127.0.0.1:5050")
            .then(response => response.json())
            .then(_read_data => {
                console.log(_read_data)
                _data.set(_read_data)
            }).catch(error => {
                console.log(error)
                return []
            })
        })
    }

    
</script>

<main>
    <h1> Todo App </h1>
    <input type="text" bind:value={taskDesc}/>
    <button on:click= { _newTask(taskDesc) }> Add Task </button>
    <div>
        {#each $_data as d}
        {#if d.Completed }
        <del> { d.Description } </del>
        <button on:click= {updateStatus(d.Description)}> Uncheck </button>
        {:else}
        <p> { d.Description } </p>
        <button on:click= {updateStatus(d.Description)}> Check </button>
        {/if}
        {/each}
    </div>
</main>