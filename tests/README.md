# Tests
These are the files used to test the sqlite database for inaccuracies

They make assumptions, such as, all translations of the bible will have the same amount of chapters and verses in those chapters

This assumption is not true, as chapters / verses are not in the original text, and each translation is free to decide how to split it up.

For the purpose of these tests however, we assume they should all match t_kjv King James Version.

If we use Crosswire https://crosswire.org/sword/modules/ModDisp.jsp?modType=Bibles as the source of the texts, we can meet these assumptions.

## Usage
These tests can be run via the include .github/workflows/ actions, or you can run ```go test``` on the command line
