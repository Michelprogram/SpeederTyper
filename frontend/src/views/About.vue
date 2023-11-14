<template>
  <div class="flex flex-col justify-center items-center font-mono">
    <p class="font-bold w-2/3 text-5xl text-center m-10 leading-tight">
      The Ultimate Typing Race Experience for Developers !
    </p>
    <div class="w-3/4">
      <p>
        Are you tired of
        <span class="italic">conventional typing games</span> that don't
        challenge your coding prowess ? SpeedyTyper is the typing race
        application designed for developers, by developer (<a
          href="https://dorian-gauron.com"
          target="_blank"
          class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
          >me</a
        >).
      </p>
    </div>
    <p class="font-bold w-2/3 text-5xl text-center m-10 leading-tight">
      How it Works
    </p>
    <div class="w-3/4 flex items-center justify-center gap-10">
      <p class="w-2/3 text-center">
        In SpeedyTyper, you'll be facing off against
        <span class="italic">fellow developers</span> in real-time typing races.
        But this isn't your average typing
        <span class="italic">competition—oh no !</span> You'll be typing out
        lines of code. Your goal ? To finish the code snippets as quickly and
        accurately as possible. Maybe one day you can beat this dude ->
      </p>

      <div class="w-44 h-44 relative rounded">
        <img
          class="rounded"
          src="https://static.frontendmasters.com/assets/teachers/theprimeagen/thumb@2x.jpg"
          alt="Rounded avatar"
        />
        <div
          class="absolute top-0 left-0 flex justify-center items-center w-full h-full bg-black rounded hover:opacity-0 duration-500"
        >
          <a
            target="_blank"
            href="https://twitter.com/ThePrimeagen"
            class="text-white"
            >The Dude</a
          >
        </div>
      </div>
    </div>
    <p class="font-bold w-2/3 text-5xl text-center m-10 leading-tight">
      The Tech Behind the Magic
    </p>
    <div class="w-[90%] flex justify-between">
      <div class="max-w-sm rounded overflow-hidden shadow-lg">
        <div class="px-6 py-4">
          <div
            class="font-bold text-xl m-2 flex flex-col items-center justify-center gap-2"
          >
            <img
              class="w-8 h-8"
              src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Objects/Mirror%20Ball.png"
              alt="Disco"
              loading="lazy"
            />
            <p>Websockets</p>
          </div>
          <p class="text-gray-700 text-base text-center">
            You wondered how your race updates in real-time ? That's the magic
            of websockets ! In SpeedyTyper, websockets ensure that every
            keystroke you make is instantly transmitted to the server and
            reflected on your opponent's screen. It's like pair programming on
            steroids !
          </p>
        </div>
        <div class="flex p-2 m-2 gap-2 items-center justify-center">
          <a
            href="https://developer.mozilla.org/fr/docs/Web/API/WebSockets_API"
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >MDN Websockets</a
          >
          <a
            href="https://pkg.go.dev/golang.org/x/net/websocket"
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >Go Websocket</a
          >
        </div>
      </div>
      <div class="max-w-sm rounded overflow-hidden shadow-lg">
        <div class="px-6 py-4">
          <div
            class="font-bold text-xl m-2 flex flex-col justify-center items-center gap-2"
          >
            <img
              class="h-8 w-8"
              src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Animals/Bear.png"
              alt="Bear"
              loading="lazy"
            />
            <p>Golang</p>
          </div>
          <p class="text-gray-700 text-base text-center">
            At the heart (and in my heart) of our server-side operations is
            Golang, the language designed by Google. Golang's speed and
            efficiency power our server, ensuring seamless communication between
            players and lightning-fast response times.
          </p>
        </div>
        <div class="flex p-2 m-2 gap-2 items-center justify-center">
          <a
            href="https://github.com/golang/go"
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >Github</a
          >

          <a
            href="https://go.dev/"
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >Go master</a
          >
        </div>
      </div>
      <div class="max-w-sm rounded overflow-hidden shadow-lg">
        <div class="px-6 py-4">
          <div
            class="font-bold text-xl m-2 flex flex-col justify-center items-center gap-2"
          >
            <img
              class="h-8 w-8"
              src="https://raw.githubusercontent.com/Tarikul-Islam-Anik/Animated-Fluent-Emojis/master/Emojis/Smilies/Robot.png"
              alt="Robot"
              loading="lazy"
            />
            <p>Vue.js</p>
          </div>
          <p class="text-gray-700 text-base text-center">
            On the frontend, I've employed the simplicity and elegance of Vue.js
            to create a user-friendly interface that keeps you focused on the
            race. With Vue, you can enjoy a smooth, reactive, and engaging
            typing experience.
          </p>
        </div>
        <div class="flex p-2 m-2 gap-2 items-center justify-center">
          <a
            href="https://vuejs.org/"
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >Vue</a
          >
          <a
            href="https://pinia.vuejs.org/"
            class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >Pinia</a
          >
        </div>
      </div>
    </div>

    <p class="font-bold w-2/3 text-5xl text-center m-10 leading-tight">
      Why Speeder Typer ?
    </p>
    <div class="w-3/4 flex items-center text-center">
      <p class="w-3/5">
        Improve your debugging skills by typing out error-ridden code snippets
        faster than your opponent can say "syntax error."
      </p>
      <p class="w-2/5">
        Connect with fellow developers worldwide and build your coding game. You
        never know when you might meet your future coding buddy or tech
        co-founder! So what are you waiting for? Join SpeedyTyper today and
        experience the thrill of a real-time coding race. Happy typing!
      </p>
    </div>

    <p class="italic text-sm m-10">
      SpeedyTyper is not responsible for any code-related epiphanies, or
      newfound love for websockets, Golang, and Vue.js. Join at your own risk!
    </p>
  </div>
</template>
<script lang="ts" setup></script>
