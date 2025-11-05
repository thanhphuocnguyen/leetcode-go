function findKthBit(n, k) {
  let sequence = '0';
  for (let i = 1; i < n && k > sequence.length; i++) {
    sequence += '1';
    for (let j = sequence.length - 2; j >= 0; j--) {
      sequence += sequence[j] === '1' ? '0' : '1';
    }
  }
  return sequence[k];
}

console.log(findKthBit(20, 1048575)); // 0