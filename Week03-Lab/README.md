# Week 03 Lab - Hashing Cryptography

 Week 03 cryptography lab tasks and outputs.

## Lab Tasks & Results

## Lab 0
![Lab 0 Result](./image-for-answer/lab0.png)

## Lab 1
![Lab 1 Result](./image-for-answer/lab1-result.png)

## Lab 2
![Lab 2 Result](./image-for-answer/lab2-result.png)

## Lab 3
![Lab 3 Result](./image-for-answer/lab3-result.png)

## Lab 4


### Process
1. **Decode URL Hash:** Used CyberChef to decode/extract the hash from the URL.
   ![CyberChef URL Decode](./lab4/cyberchef-encode-url.png)

2. **Crack MD5 Hash:** Used an MD5 lookup/reverse tool to crack the extracted hash, resulting in `meowmeow`.
   ![MD5 Reverse Crack](./lab4/crack-md5-hash.png)

3. **Format & Verify Answer:** Formatted the flag as `cryptoCTF{meowmeow}` and verified it against the given regex instruction `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`, confirming a match.
   ![Regex Verification Match](./image-for-answer/lab4-result-verify.png)

