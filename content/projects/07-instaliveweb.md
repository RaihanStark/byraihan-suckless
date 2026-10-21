---
title: InstaLiveWeb: Instagram Live Dashboard
shortTitle: InstaLiveWeb
slug: instaliveweb
homeDesc: Instagram Live RTMP dashboard
listDesc: third-party Instagram Live RTMP dashboard for OBS Studio
meta: Client: PT Ajita Indonesia
description: third-party Instagram Live RTMP dashboard for OBS Studio
---

<img src="/images/projects/instaliveweb/instaliveweb.png" alt="InstaLiveWeb Dashboard" loading="lazy" decoding="async">

      <h2>Introduction</h2>
      <p>
        An Instagram Live Third-Party Platform to provides you with a RTMP server and stream key for streaming using software like OBS-Studio.
      </p>

      <h2>Problems</h2>
      <p>
        Instagram did not provide RTMP for user to stream outside Instagram Client. This limitation prevented content creators from:
      </p>
      <ul>
        <li>Using professional streaming software</li>
        <li>Streaming from desktop computers</li>
        <li>Utilizing advanced streaming features like overlays, multiple cameras, and screen sharing</li>
        <li>Managing live comments and viewers from a central dashboard</li>
      </ul>

      <h2>Solution</h2>
      <p>
        We reverse engineered the Instagram API to get the RTMP from the Instagram Server, and we created a dashboard for it so user can easily use and manage it. This solution enables professional content creators to leverage their existing streaming setup while reaching their Instagram audience.
      </p>

      <h2>Key Features</h2>

      <h3>Streaming Capabilities</h3>
      <ul>
        <li><strong>RTMP Server Generation</strong>: Automatically generates RTMP server URL and stream key</li>
        <li><strong>OBS Integration</strong>: Seamless integration with OBS-Studio and other streaming software</li>
        <li><strong>Broadcast Control</strong>: Start and stop broadcasts directly from the dashboard</li>
      </ul>

      <h3>Live Management</h3>
      <ul>
        <li><strong>Live Comments</strong>: Real-time display of viewer comments</li>
        <li><strong>Live Viewers</strong>: Track current viewer count and engagement</li>
        <li><strong>Instagram Login</strong>: Secure authentication with your Instagram account</li>
        <li><strong>OTP Handler</strong>: Automatic handling of Instagram's two-factor authentication</li>
      </ul>

      <h3>User Experience</h3>
      <ul>
        <li><strong>Intuitive Dashboard</strong>: Clean and easy-to-use interface</li>
        <li><strong>Real-time Updates</strong>: Live updates for comments and viewer metrics</li>
        <li><strong>Responsive Design</strong>: Works on desktop and mobile devices</li>
      </ul>

      <h2>Technology Stack</h2>

      <h3>Backend</h3>
      <ul>
        <li><strong>Python</strong> - Core programming language</li>
        <li><strong>Flask</strong> - Web framework</li>
        <li><strong>Flask-WTF</strong> - Form handling and validation</li>
        <li><strong>Flask-Session</strong> - Session management</li>
      </ul>

      <h3>Frontend</h3>
      <ul>
        <li><strong>JavaScript</strong> - Client-side functionality</li>
        <li><strong>HTML/SCSS</strong> - Markup and styling</li>
        <li><strong>jQuery</strong> - DOM manipulation</li>
        <li><strong>Ajax</strong> - Asynchronous requests</li>
      </ul>

      <h3>Infrastructure</h3>
      <ul>
        <li><strong>Heroku</strong> - Cloud hosting platform</li>
        <li><strong>Instagram Private API</strong> - Reverse-engineered Instagram endpoints</li>
      </ul>

      <h2>Technical Implementation</h2>
      <p>The project involved several technical challenges:</p>
      <ol>
        <li><strong>API Reverse Engineering</strong>: Carefully analyzing Instagram's private API to understand RTMP stream generation</li>
        <li><strong>Authentication Flow</strong>: Implementing secure Instagram login with OTP support</li>
        <li><strong>Real-time Communication</strong>: Building WebSocket connections for live comments and viewer updates</li>
        <li><strong>Stream Management</strong>: Handling RTMP server allocation and stream lifecycle</li>
      </ol>

      <h2>Screenshots</h2>
      <h3>Dashboard Interface</h3>
      <img src="/images/projects/instaliveweb/instaliveweb-2.png" alt="InstaLiveWeb Dashboard Interface" loading="lazy" decoding="async">

      <h2>Impact</h2>
      <p>InstaLiveWeb revolutionized how content creators stream to Instagram by:</p>
      <ul>
        <li>Enabling professional-quality streams with advanced software features</li>
        <li>Providing better engagement tools through the live comment dashboard</li>
        <li>Allowing desktop streaming for the first time</li>
        <li>Supporting multi-camera and screen sharing capabilities</li>
      </ul>
      <p>
        The platform served thousands of content creators who wanted to elevate their Instagram Live experience beyond mobile-only limitations.
      </p>
