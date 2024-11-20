import { animate } from 'motion'

// Add this class to an element containing text for Sandman to use as letter particles
export const SANDMAN_TEXT_CLASS = 'sandman-text'

const LETTER_CLASS = 'sandman-letter',
  INITIAL_OFFSET_TOP = 'sandmanTop',
  INITIAL_OFFSET_BOTTOM = 'sandmanBottom',
  INITIAL_OFFSET_LEFT = 'sandmanLeft',
  INITIAL_OFFSET_RIGHT = 'sandmanRight',
  CURRENT_TOP = 'sandmanCurrTop'

export class Sandman {
  // This is the container element we use to search for
  // targets containing .sandman-text classes,
  // and to position letters.
  private container: HTMLElement

  private initialized: boolean
  private maxDropScrollTop: number

  private originalValues: Map<Element, string>

  constructor(container: HTMLElement) {
    this.container = container
    this.initialized = false
    this.maxDropScrollTop = 0
    this.originalValues = new Map()
  }

  // Initializes Sandman.
  init() {
    if (this.initialized) {
      return
    }

    this.replaceLettersWithSpans()

    this.initialized = true
  }

  // De-initialize Sandman, and return to normal state.
  destroy() {
    for (const [el, html] of this.originalValues.entries()) {
      el.innerHTML = html
    }

    this.initialized = false
    this.maxDropScrollTop = 0
  }

  // Resets animations by destroying and initializing once again.
  reset() {
    this.destroy()
    this.init()
  }

  // Call this method whenever the container has been resized IF Sandman has been initialized.
  // Sandman will recalculate each letter's initial offset positions.
  onContainerResize() {
    const spans = this.container.getElementsByClassName(LETTER_CLASS)
    for (const s of spans) {
      const span = s as HTMLElement
      const spanBounds = span.getBoundingClientRect()

      const deltaY = spanBounds.top - Number(span.dataset[CURRENT_TOP])

      span.dataset[INITIAL_OFFSET_TOP] = (
        Number(span.dataset[INITIAL_OFFSET_TOP]) + deltaY
      ).toString()
      span.dataset[INITIAL_OFFSET_BOTTOM] = (
        Number(span.dataset[INITIAL_OFFSET_BOTTOM]) + deltaY
      ).toString()
      span.dataset[CURRENT_TOP] = spanBounds.top.toString()
    }
  }

  async setText(
    text: HTMLElement,
    letterDelay: number = 0,
    duration?: number,
    bounce?: number,
  ): Promise<Sandman> {
    if (!this.initialized) {
      return this
    }

    return new Promise<Sandman>((resolve) => {
      const spans = text.getElementsByClassName(LETTER_CLASS)
      let animFinishCount = 0
      let delay = 0

      for (const s of spans) {
        const span = s as HTMLElement
        span.style.visibility = ''
        animate(
          span,
          {
            x: 0,
            y: 0,
            rotate: 0,
          },
          {
            type: 'spring',
            ease: 'easeInOut',
            duration,
            bounce,
            delay,
          },
        ).then(() => {
          span.dataset[CURRENT_TOP] = span.getBoundingClientRect().top.toString()

          animFinishCount++

          if (animFinishCount === spans.length) {
            resolve(this)
          }
        })

        delay += letterDelay
      }
    })
  }

  // TODO: Specify size and position, vector, speed
  async collect(speed: number, maxLetters?: number): Promise<Sandman> {
    if (!this.initialized) {
      return this
    }

    const containerBounds = this.container.getBoundingClientRect()
    const zoneWidth = containerBounds.width * 0.2
    const zoneHeight = containerBounds.height * 0.2
    const zoneLeft = containerBounds.left + containerBounds.width / 2 - zoneWidth / 2
    const zoneRight = zoneLeft + zoneWidth
    const zoneTop = containerBounds.top + containerBounds.height / 2 - zoneHeight / 2
    const zoneBottom = zoneTop + zoneHeight
    let animFinishCount = 0

    return new Promise<Sandman>((resolve) => {
      const observer = new IntersectionObserver((entries) => {
        for (const entry of entries) {
          const span = entry.target as HTMLElement
          const spanBounds = entry.boundingClientRect
          const hide = maxLetters && !spansToDisplay.includes(span)

          const deltaX = this.randomNumber(
            zoneLeft - (containerBounds.left + Number(span.dataset[INITIAL_OFFSET_LEFT])),
            zoneRight - (containerBounds.left + Number(span.dataset[INITIAL_OFFSET_RIGHT])),
          )

          const deltaY = this.randomNumber(
            zoneTop - (containerBounds.top + Number(span.dataset[INITIAL_OFFSET_TOP])),
            zoneBottom - (containerBounds.top + Number(span.dataset[INITIAL_OFFSET_BOTTOM])),
          )

          if (hide) {
            span.style.visibility = 'hidden'
          }

          animate(
            entry.target,
            {
              x: deltaX,
              y: deltaY,
            },
            {
              type: 'spring',
              bounce: 0.3,
              delay: this.randomNumber(0, 0.3),
              duration: hide
                ? 0
                : Math.sqrt(
                    this.calcDistance(
                      { x: spanBounds.left, y: spanBounds.top },
                      { x: zoneLeft, y: zoneTop },
                    ),
                  ) / speed,
              ease: 'easeInOut',
            },
          ).then(() => {
            span.dataset[CURRENT_TOP] = this.calcNewPosition(span, containerBounds, {
              y: deltaY,
            }).y.toString()

            animFinishCount++

            if (animFinishCount === spans.length) {
              resolve(this)
            }
          })

          observer.unobserve(entry.target)
        }
      })

      const spans = this.container.getElementsByClassName(LETTER_CLASS)
      const spansToDisplay: Element[] = []
      if (maxLetters) {
        for (let i = 0; i < maxLetters; i++) {
          spansToDisplay.push(this.randomFromArray([...spans]))
        }
      }

      for (const span of spans) {
        observer.observe(span)
      }
    })
  }

  // Drops all text to the bottom of the container.
  async drop(speed: number): Promise<Sandman> {
    if (!this.initialized) {
      // TODO: Throw error?
      return this
    }

    // Prevent text from returning top, only drop to bottom
    if (this.maxDropScrollTop !== 0 && this.container.scrollTop <= this.maxDropScrollTop) {
      return this
    }
    this.maxDropScrollTop = this.container.scrollTop

    const containerBounds = this.container.getBoundingClientRect()
    let animFinishCount = 0

    return new Promise<Sandman>((resolve) => {
      const observer = new IntersectionObserver((entries) => {
        for (const entry of entries) {
          const span = entry.target as HTMLElement
          const spanBounds = entry.boundingClientRect

          // Ignore elements already lower than container's lower bound
          if (spanBounds.bottom >= containerBounds.bottom) {
            animFinishCount++
          } else {
            const deltaX = this.randomNumber(-50, 50)

            const deltaY =
              containerBounds.bottom +
              this.maxDropScrollTop -
              (containerBounds.top + Number(span.dataset[INITIAL_OFFSET_BOTTOM]))

            animate(
              entry.target,
              {
                x: deltaX,
                y: deltaY,
                rotate: this.randomNumber(-90, 90),
              },
              {
                duration: Math.sqrt(
                  this.calcDistance(
                    { x: spanBounds.left, y: spanBounds.top },
                    { x: deltaX, y: deltaY },
                  ) / speed,
                ),
                delay: this.randomNumber(0, 0.3),
                ease: [0.38, 0.32, 0.9, 0.44],
              },
            ).then(() => {
              span.dataset[CURRENT_TOP] = this.calcNewPosition(span, containerBounds, {
                y: deltaY,
              }).y.toString()

              animFinishCount++

              if (animFinishCount === spans.length) {
                resolve(this)
              }
            })
          }

          observer.unobserve(entry.target)
        }
      })

      const spans = this.container.getElementsByClassName(LETTER_CLASS)
      for (const span of spans) {
        observer.observe(span)
      }
    })
  }

  // Replaces each letter inside target (elements with .sandman-text)
  // with a <span> encompassing the letter.
  private replaceLettersWithSpans() {
    const containerBounds = this.container.getBoundingClientRect()
    this.container.dataset[INITIAL_OFFSET_TOP] = containerBounds.top.toString()
    this.container.dataset[INITIAL_OFFSET_BOTTOM] = containerBounds.bottom.toString()
    this.container.dataset[INITIAL_OFFSET_LEFT] = containerBounds.left.toString()
    this.container.dataset[INITIAL_OFFSET_RIGHT] = containerBounds.left.toString()

    const targetElements = this.container.getElementsByClassName(SANDMAN_TEXT_CLASS)

    for (const el of targetElements) {
      const letters = el.textContent?.split('') ?? []

      this.originalValues.set(el, el.innerHTML)
      el.innerHTML = ''

      const spans: HTMLSpanElement[] = []
      for (const letter of letters) {
        const span = document.createElement('span')
        span.classList.add(LETTER_CLASS)
        span.textContent = letter

        // We add a white-space: pre css to spaces
        // So their widths don't become 0.
        if (letter === ' ') {
          span.style.whiteSpace = 'pre'
        }

        span.style.display = 'inline-block'

        el.appendChild(span)

        const bounds = span.getBoundingClientRect()
        span.dataset[INITIAL_OFFSET_TOP] = (bounds.top - containerBounds.top).toString()
        span.dataset[INITIAL_OFFSET_BOTTOM] = (bounds.bottom - containerBounds.top).toString()
        span.dataset[INITIAL_OFFSET_LEFT] = (bounds.left - containerBounds.left).toString()
        span.dataset[INITIAL_OFFSET_RIGHT] = (bounds.right - containerBounds.left).toString()
        span.dataset[CURRENT_TOP] = bounds.top.toString()

        spans.push(span)
      }
    }
  }

  // Generates random number between min (inclusive) and max (inclusive).
  private randomNumber(min: number, max: number) {
    return Math.random() * (max - min + 1) + min
  }

  private randomFromArray<T>(arr: Array<T>) {
    return arr[Math.floor(Math.random() * arr.length)]
  }

  private calcDistance<Point2d extends { x: number; y: number }>(
    from: Point2d,
    to: Point2d,
  ): number {
    return Math.min(10, Math.pow(to.x - from.x, 2) + Math.pow(to.x - from.x, 2))
  }

  private calcNewPosition(el: HTMLElement, containerBounds: DOMRect, delta: { y: number }) {
    return { y: containerBounds.top + Number(el.dataset[INITIAL_OFFSET_TOP]) + delta.y }
  }
}
