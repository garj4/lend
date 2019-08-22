## Description

This project creates a CLI tool (`lend`) that allows the user to easily track how much money they owe others.

Lender allows you to easily record the amount of money that changes hands among friends.
Not everyone has the ability to make an immediate reimbursement, and sometimes money changes hands so frequently that it's simply not worth it to constantly make bank transfers.

## Installation

From the root project directory:

```
go install
```

## Usage

The CLI is invoked with `lend`, followed by the specific command to run. The `--help` flag will print help information for any command.

## Structure

This project uses [Cobra](https://github.com/spf13/cobra), a CLI development framework for golang. The basis of the project was generated with the [cobra generator](https://github.com/spf13/cobra/blob/master/cobra/README.md).

## Licensing

Copyright © 2019 Garrett Wezniak <gaarj4@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
