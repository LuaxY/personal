<template>
  <div id="works">
    <div class="container mx-auto">
      <div class="grid grid-cols-2 gap-32">
        <div class="grid grid-cols-1 gap-20">
          <div class="text">
            <h3>Latest Works</h3>
            <p>An overview on what I have build</p>
          </div>
          <div class="block nitro-encoder">
            <h4>Nitro Video Encoder</h4>
            <p>Nitro Encoder is a powerful transcoder of video into streaming format in a distributed manner. It splits movies into small chunks to allow parallel encoding in the cloud at a reduced price of preemptible GCP instances.</p>
          </div>
          <div class="text">
            <h3>All Projects</h3>
            <p>If you want to know more, contact me</p>
          </div>
        </div>
        <div class="grid grid-cols-1 gap-32">
          <div class="block ubisoft-monitoring">
            <h4>Ubisoft Monitoring Platform</h4>
            <p>Implementation of a Monitoring-as-a-Service system, to monitor Ubisoft's entire network infrastructure and its hundreds of thousands of daily players. MonASS uses the power of the multi cloud to be highly available, fault tolerant and recoverable in case of disasters.</p>
          </div>
          <div class="block dofus-anti-cheat">
            <h4>Dofus Anti Cheat</h4>
            <p>Realization of an anti cheating system for the popular game Dofus in France. It aims to fight against bots using several techniques such as behavior detection, machine learning, crypto whiteboxing, and security by obscurity (blackboxing). It also allows the detection of illegal gold sellers and so far no one has been able to bypass this system.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Works"
  }
</script>