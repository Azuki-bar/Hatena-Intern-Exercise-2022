export function initializeFizzBuzz(
  list: HTMLOListElement,
  button: HTMLButtonElement
): void {
  console.debug("initializeFizzBuzz", list, button);
  // Implement Me!
  const fizzbuzz = new FizzBuzz();
  button.addEventListener('click', () => {
    const newElem = document.createElement('li')
    newElem.textContent = fizzbuzz.next()
    list.appendChild(newElem)
  })
}

interface IncrementValueValue {
  next(): string;
}

class FizzBuzz implements IncrementValueValue {
  private now: number;
  constructor() {
    this.now = 0;
  }
  next(): string {
    this.now++;
    if (this.now % 15 == 0) {
      return "FizzBuzz";
    } else if (this.now % 3 == 0) {
      return "Fizz";
    } else if (this.now % 5 == 0) {
      return "Buzz";
    }
    return this.now.toString();
  }
}
