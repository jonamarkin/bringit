<template>
  <article class="card cashflow-card">
    <div class="card-header">
      <div>
        <h2>Cash Flow</h2>
      </div>
      <button class="link-button" type="button">View more</button>
    </div>

    <div class="cashflow-summary">
      <strong>$95,193.00</strong>
      <p>vs 2,293.00 last month</p>
    </div>

    <div class="cashflow-stack">
      <div class="cashflow-block">
        <div class="cashflow-topline">
          <strong>$95,193.00</strong>
          <p>vs 2,293.00 last month</p>
        </div>
        <div class="progress-bar green-bar">
          <span></span>
        </div>
        <div class="progress-meta">
          <span>60% money in</span>
        </div>
      </div>

      <div class="cashflow-block">
        <div class="cashflow-topline">
          <strong>$95,193.00</strong>
          <p>vs 2,293.00 last month</p>
        </div>
        <div class="progress-bar red-bar">
          <span></span>
        </div>
        <div class="progress-meta">
          <span>40% money out</span>
        </div>
      </div>
    </div>

    <div class="compact-filters lower-filters">
      <button class="compact-pill active" type="button">1W</button>
      <button class="compact-pill" type="button">1M</button>
      <button class="compact-pill" type="button">6M</button>
      <button class="compact-pill" type="button">1Y</button>
      <button class="compact-pill with-icon" type="button">
        <svg viewBox="0 0 16 16" aria-hidden="true">
          <path d="M3.2 4.3H12.8"></path>
          <path d="M5.4 8H10.6"></path>
          <path d="M7 11.7H9"></path>
        </svg>
        <span>Filter</span>
      </button>
    </div>
  </article>
</template>
