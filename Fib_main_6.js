<script>
    let n = 9;
     
    function fib(n) {
    if (n <= 1)
        return n;
    return fib(n-1) + fib(n-2);
}
  
    document.write(fib(n));
</script>
