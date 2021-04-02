List<int> fibonacci(int n) {
  int last = 1;
  int last2 = 0;
  return List<int>.generate(n, (int i) {
    if (i < 2) return i;
    int curr = last + last2;
    last2 = last;
    last = curr;
    return curr;
  });
}
