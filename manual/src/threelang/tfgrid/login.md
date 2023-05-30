# Login Namespace

- To login with your credentials on the tfgrid, use the login namespace
- The login action is special, it does not have sub operations
- action name: !!tfgrid.login
- parameters:
  - mnemonic [required]
  - network [optional]
    - a string in ['dev', 'qa', 'test', 'main']
    - if not provided, defaults to 'main'
