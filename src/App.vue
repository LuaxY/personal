<template>
  <div id="app">
    <div class="header">
      <div class="container mx-auto">
        <div class="name">
          <b>Yann</b> Guineau
        </div>
        <div class="links">
          <a href="https://twitter.com/Luax0" target="_blank" class="link"><i class="fab fa-twitter"></i> Twitter</a>
          <a href="https://github.com/LuaxY" target="_blank" class="link"><i class="fab fa-github"></i> GitHub</a>
          <a href="https://keybase.io/luax" target="_blank" class="link"><i class="fab fa-keybase"></i> Keybase</a>
          <a href="https://www.linkedin.com/in/yann-guineau/" target="_blank" class="link"><i class="fab fa-linkedin"></i> LinkedIn</a>
          <a href="https://calendly.com/yann-guineau/30min" target="_blank" class="link"><i class="fas fa-calendar-alt"></i> Meeting</a>
          <a href="https://www.linkedin.com/in/yann-guineau/" target="_blank" class="btn btn-yellow">Hire me</a>
        </div>
      </div>
    </div>

    <router-view/>

    <div id="contact" class="footer">
      <div class="container mx-auto">
        <div class="grid grid-cols-3 gap-32">
          <div>
            <h3>Got a project?<br>Let's talk.</h3>
            <p>If you are interested to colaborate with me, don't hesitate to send me an email.</p>
            <div class="email">
              <a href="mailto:hello@voidmx.net?subject=I have an awesome idea!!"><i class="fas fa-envelope"></i> hello@voidmx.net</a>
            </div>
          </div>
          <div class="col-span-2 networks flex-end">
            <div class="inline-block">
              <a href="https://twitter.com/Luax0" target="_blank" class="link"><i class="fab fa-twitter"></i></a>
              <a href="https://github.com/LuaxY" target="_blank" class="link"><i class="fab fa-github"></i></a>
              <a href="https://keybase.io/luax" target="_blank" class="link"><i class="fab fa-keybase"></i></a>
              <a href="https://www.linkedin.com/in/yann-guineau/" target="_blank" class="link"><i class="fab fa-linkedin"></i></a>
              <a href="https://calendly.com/yann-guineau/30min" target="_blank" class="link"><i class="fas fa-calendar-alt"></i></a>
            </div>
            <div class="code"><a href="https://github.com/LuaxY/personal" target="_blank"><i class="far fa-code"></i> View Code Source</a></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
