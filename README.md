<h1>
    PuppetLang 🎩
</h1>
PuppetLang is a domain-specific programming language designed specifically for writing scripts that automate web browsers in a headless environment, like Puppeteer. It provides a specialized syntax and features tailored for browser automation tasks, making it easier to develop and maintain automated browser scripts compared to using general-purpose languages.

<h1>
    Installation 🧨
</h1>
Clone this repo and build the project:

```
git clone https://github.com/RaniGiterman/PuppetLang.git PuppetLang
cd PuppetLang
go build -o pu .
sudo mkdir -p /usr/local/pu/bin
sudo cp pu /usr/local/pu/bin
export PATH=$PATH:/usr/local/pu/bin
pu -h
```

<h1>
    Get Started 🎯
</h1>
Now that you have PuppetLang installed, let's write some example code! <br>
Here are the available commands: <br>
<table>
  <tr>
    <th>Command</th>
    <th>Description</th>
  </tr>
  <tr>
    <td>url "&lt;ADDRESS&gt;"</td>
    <td>Travels to requested address.</td>
  </tr>
  <tr>
    <td>click "&lt;SELECTOR&gt;"</td>
    <td>Clicks the requested element, using js selector.</td>
  </tr>
  <tr>
    <td>write "&lt;SELECTOR&gt;" "&lt;TEXT&gt;"</td>
    <td>Clicks the requested element, using js selector.</td>
  </tr>
  <tr>
    <td>screenshot "&lt;PATH&gt;"</td>
    <td>Takes a full screenshot of the screen and saves the png picture to the requested path.</td>
  </tr>
</table>

Now that you are familar with the commands, let's put them all into action. <br>
The following script uses the commands you've learned to travel to a site, and take 3 screenshots. <br>

```
url "https://rg-playground.vercel.app/"
screenshot "a.png"

click "#run_code"
screenshot "b.png"

write "#editor" "הדפסה(1)"
click "#run_code"
screenshot "c.png"
```

Use the example `script.pu` file included in the repo, or paste the script into a file of your own. Finally, run the code:

```
pu -f script.pu -t 30s
```

The above command targets the pu exec to run the file `script.pu` using the `-f` flag, and sets a timeout duration of 30 seconds. <br>
You should see 3 screenshots saved on your current directory, named `a.png`, `b.png` and `c.png`. Good Luck!
