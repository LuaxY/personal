<template>
  <div class="home">
    <Me/>
    <About/>
    <Companies/>
    <Works/>
    <div id="technologies"></div> <!-- TODO -->
  </div>
</template>

<script>
  import Me from '../components/Me.vue'
  import About from '../components/About.vue'
  import Companies from '../components/Companies.vue'
  import Works from '../components/Works.vue'

  export default {
    name: 'Home',
    components: {
      Me,
      About,
      Companies,
      Works
    }
  }
</script>
