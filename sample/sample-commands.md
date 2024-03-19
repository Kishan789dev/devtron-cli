

# Setup instruction
1. Download and unpack OS specific binary from git repo.
2. add alias devtctl to bashrc/zshrc to the downloaded path  
    - copy the path of downloaded binary  
    - `vim ~/.zshrc or vim ~/.bashrc` (check the default set by running echo $0 )  
    - add line  `alias="Your binary path/devtctl"`  
    - save and run `source ~/.bashrc` or `source ~/.zshrc` or restart your terminal  
4. create devtctl.env `vim devtctl.env` file at the path of binary with the following keys
5.     SERVER_URL= "https://subhas.devtron.info/"
       API_TOKEN="your API token"  
    

7. enable developer options on mac for terminal if not already enabled (System preferences -> security&Privacy -> Privacy tab -> Developer tools -> enable terminal here)  

# Help section
    devtctl --help  
    devtctl getCI --help  
    devtctl applyCI --help  

# Sample commands for getCI and applyCI

## To download CI config passing criteria through flags
    devtctl getCI --pipelineIds="1 2 3"  
    devtctl getCI --appIds="1 2 3"  
    devtctl getCI --appNames="app1 app2"  
    devtctl getCI --projectNames="project1 project2"  
    devtctl getCI --envNames="env1 env2"  

All these flags can be used in any combinations. In the sample inputManifest files, there are additional, more complex criteria that users can utilize: `includesAppNames` and `excludesAppNames`. These criteria use regular expressions to search for and filter apps based on their app names.

Let's explain these criteria in a more straightforward way for users to easily understand:

`includesAppNames`: This criteria allows you to specify a regular expression to filter apps based on their app names. If you define includesAppNames, it works similarly to using the --appNames flag via the command-line interface (CLI). By setting includesAppNames, you will receive all the apps whose app names match the provided regular expression.

`excludesAppNames`: In addition to the `includesAppNames`, this criteria allows you to further refine the selection by excluding certain apps. If you use excludesAppNames along with all other criteria, the included apps from `includesAppNames` plus all other criteria mentioned are evaluated first. Then, based on the excludesAppNames regex, any apps with matching app names are excluded from the final result.

To specify the format of output use flag  
    
    --format=json   
    --format=yaml    
(yaml is default)  
To get the output manifest in a file use flag  

    -o "./output.yaml"

## To download CI config passing criteria through manifest file 
    use flag -p "./inputmanifest.yaml"
Provide metadata type as YAML_DOWNLOAD or JSON_DOWNLOAD


## To patch preCI-preCD for existing CI piplines

Update the downloaded manifest by following the steps  
1. change the metadata type to PATCH  
2. one entry in payload is for one ci-pipeline associated to it.  
3. change the criteria as per your needs for which the particular payload will be used to path  
4. update the pre-ci/post-ci steps as necessary 
5. save the file and run the command  
`
    devtctl applyCI -p ./inputmanifest.yaml
`  
`
    devtctl applyCI -p ./inputmanifest.json
`   
(both yaml/json type suported in input)

to disable interactive prompts and reply yes to all use  
`flag   --allYes  `  
to patch with empty pre or post CI stage(which will lead to deletion of current configuration) use   
`
flag --overrideEmptySteps
`











