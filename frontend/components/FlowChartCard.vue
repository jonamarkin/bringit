<template>
  <article class="card chart-card">
    <div class="card-header">
      <div>
        <h2>Transactions History</h2>
        <p>All general information appears in this page.</p>
      </div>
      <div class="compact-filters">
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
    </div>

    <div class="chart-area">
      <div class="chart-y-axis">
        <span>40K</span>
        <span>35K</span>
        <span>30K</span>
        <span>25K</span>
        <span>20K</span>
        <span>15K</span>
        <span>10K</span>
        <span>0K</span>
      </div>

      <div class="chart-canvas">
        <svg viewBox="0 0 640 270" aria-label="Transactions history chart">
          <defs>
            <linearGradient id="greenFade" x1="0" x2="0" y1="0" y2="1">
              <stop offset="0%" stop-color="#2fe28d" stop-opacity="0.24"></stop>
              <stop offset="100%" stop-color="#2fe28d" stop-opacity="0"></stop>
            </linearGradient>
            <linearGradient id="redFade" x1="0" x2="0" y1="0" y2="1">
              <stop offset="0%" stop-color="#ff8b8b" stop-opacity="0.12"></stop>
              <stop offset="100%" stop-color="#ff8b8b" stop-opacity="0"></stop>
            </linearGradient>
          </defs>

          <g class="grid-lines">
            <line x1="26" y1="24" x2="624" y2="24"></line>
            <line x1="26" y1="55" x2="624" y2="55"></line>
            <line x1="26" y1="86" x2="624" y2="86"></line>
            <line x1="26" y1="117" x2="624" y2="117"></line>
            <line x1="26" y1="148" x2="624" y2="148"></line>
            <line x1="26" y1="179" x2="624" y2="179"></line>
            <line x1="26" y1="210" x2="624" y2="210"></line>
            <line x1="26" y1="241" x2="624" y2="241"></line>
          </g>

          <path class="area green-area" d="M26 218C54 216 64 138 95 103C121 73 151 85 176 94C203 103 236 97 266 95C293 93 314 97 343 89C366 82 392 74 425 75C449 75 474 82 497 96C520 111 549 118 570 138C592 160 610 162 624 160L624 242L26 242Z"></path>
          <path class="area red-area" d="M26 232C52 225 70 164 95 136C121 109 144 118 168 131C196 147 219 142 246 140C277 138 299 141 329 138C358 135 381 138 409 151C439 164 459 175 489 175C516 175 533 165 561 171C587 177 608 182 624 183L624 242L26 242Z"></path>

          <path class="curve green-curve" d="M26 218C54 216 64 138 95 103C121 73 151 85 176 94C203 103 236 97 266 95C293 93 314 97 343 89C366 82 392 74 425 75C449 75 474 82 497 96C520 111 549 118 570 138C592 160 610 162 624 160"></path>
          <path class="curve red-curve" d="M26 232C52 225 70 164 95 136C121 109 144 118 168 131C196 147 219 142 246 140C277 138 299 141 329 138C358 135 381 138 409 151C439 164 459 175 489 175C516 175 533 165 561 171C587 177 608 182 624 183"></path>

          <line class="focus-line" x1="249" y1="14" x2="249" y2="242"></line>
          <circle class="focus-point green-point" cx="249" cy="95" r="5"></circle>
          <circle class="focus-point red-point" cx="249" cy="140" r="5"></circle>
        </svg>

        <div class="chart-tooltip">
          <strong>May 2024</strong>
          <div>
            <span><i class="dot green"></i>Income</span>
            <span>$12,000</span>
          </div>
          <div>
            <span><i class="dot red"></i>Expense</span>
            <span>$10,000</span>
          </div>
        </div>

        <div class="chart-months">
          <span>Jan</span>
          <span>Feb</span>
          <span>Mar</span>
          <span>Apr</span>
          <span>May</span>
          <span>Jun</span>
          <span>Jul</span>
          <span>Aug</span>
          <span>Sep</span>
          <span>Oct</span>
          <span>Nov</span>
          <span>Dec</span>
        </div>
      </div>
    </div>
  </article>
</template>
