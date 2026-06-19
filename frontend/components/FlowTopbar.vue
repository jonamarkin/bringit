<template>
  <header class="topbar">
    <div class="mobile-topbar">
      <div class="mobile-brand">
        <span class="brand-mark" aria-hidden="true"></span>
        <span class="brand-name">BringIt</span>
      </div>

      <div class="mobile-topbar-actions">
        <button class="icon-button mobile-only-button" type="button" aria-label="Notifications">
          <svg viewBox="0 0 16 16" aria-hidden="true">
            <path d="M8 2.8C6 2.8 4.6 4.2 4.6 6.2V8.2L3.6 10.1H12.4L11.4 8.2V6.2C11.4 4.2 10 2.8 8 2.8Z"></path>
            <path d="M6.5 11.4C6.8 12.3 7.2 12.7 8 12.7C8.8 12.7 9.2 12.3 9.5 11.4"></path>
          </svg>
          <span class="alert-dot"></span>
        </button>
      </div>
    </div>

    <div class="topbar-main">
      <div class="search-box">
        <svg viewBox="0 0 16 16" aria-hidden="true">
          <circle cx="7" cy="7" r="4.4"></circle>
          <path d="M10.5 10.5L13.2 13.2"></path>
        </svg>
        <span>Search</span>
        <div class="search-shortcut">
          <span>⌘</span>
          <span>F</span>
        </div>
      </div>

      <div class="topbar-actions">
        <button class="icon-button" type="button" aria-label="Notifications">
          <svg viewBox="0 0 16 16" aria-hidden="true">
            <path d="M8 2.8C6 2.8 4.6 4.2 4.6 6.2V8.2L3.6 10.1H12.4L11.4 8.2V6.2C11.4 4.2 10 2.8 8 2.8Z"></path>
            <path d="M6.5 11.4C6.8 12.3 7.2 12.7 8 12.7C8.8 12.7 9.2 12.3 9.5 11.4"></path>
          </svg>
          <span class="alert-dot"></span>
        </button>

        <button class="profile-chip" type="button">
          <span class="avatar-ring">
            <span class="avatar-face"></span>
          </span>
          <span class="profile-name">Mikolin Slavana</span>
          <svg class="chevron" viewBox="0 0 16 16" aria-hidden="true">
            <path d="M4.5 6.5L8 10L11.5 6.5"></path>
          </svg>
        </button>
      </div>
    </div>
  </header>
</template>
