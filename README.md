# ggrep
:mag: CLI utility to quickly browse GitHub repositories

## How to setup?
### A. Compile by yourself
1. Clone this repository on your local machine
2. `cd` to cloned repository and run `make build` (`Golang` compiler required)
3. Now binary `ggrep` in `/usr/local/bin/` created

### B. Pull image and run via Docker
1. Run to pull container and run shell inside it
    ```shell
    docker run -it --rm ggrep
    ```
## How to use?
### Launch
```shell
./ggrep [-lang=<programming_language>] <token1> ... <tokenX>
```
### Example
```shell
~ # ./ggrep -lang=C# bmstu cg
1. [Winterpuma/bmstu_CG]
Link: https://github.com/Winterpuma/bmstu_CG
Language: C#

2. [Winterpuma/bmstu_CG_CP]
Link: https://github.com/Winterpuma/bmstu_CG_CP
Language: C#

3. [Sunshine-ki/BMSTU5_CG_CP]
Link: https://github.com/Sunshine-ki/BMSTU5_CG_CP
Language: C#

4. [DeaLoic/bmstu-course-CG]
Link: https://github.com/DeaLoic/bmstu-course-CG
Language: C#

5. [jokerva1010123/BMSTU-5-cem-CGCW]
Link: https://github.com/jokerva1010123/BMSTU-5-cem-CGCW
Language: C#

6. [GrishinEgor/BMSTU_IU7_CourseProject_CG]
Link: https://github.com/GrishinEgor/BMSTU_IU7_CourseProject_CG
Language: C#

Total: 6
```
