# Adding Texts from ESword Website Source

If you are adding a new text to this repo, be careful! Its the Bible. Mistakes must be avoided. On the ESWord Project there are a variety of languages and translations. Each one will no doubt have it's own 
unique issues to overcome in conversion. 

Be sure you have the dependencies for this project installed:

```
pip install mysql-connector-python
pip install future
pip install pysword
pip install pyyaml
```

## Step One: Find the translation...

Go to the main downloads page: https://crosswire.org/sword/modules/ModDisp.jsp?modType=Bibles&form=MG0AV3

Find the translation you are seeking under the relevant language code. 

In this case, we are doing **Westminster Leningrad Codex** under the language code **hbo** (ancient Hebrew). 

## Step Two: Check the info page...

Next we must check the info page to ensure it has a permissive license. This would be **Public Domain** or any other permissive license, free for distribution. 

## Step Three: Create directory and add README.md

Now you will go to the `sources` directory and create the language code if it does not yet exist (in this case **hbo**)

Next, you must create the bible version directory. In this case it will be WLC. 

> **Note:** The directory must match this convention: `sources/<language>/<translation>/`

...which in this case is `sources/hbo/WLC/`

This is a very sensitive convention, where upper and lower case matters. You will see why soon.

Create a new file `README.md` and add the following info: Book Name and License, like so:

```
# WLC: Westminster Leningrad Codex

**License:** Public Domain
```

This README.md should be placed in the `<translation>/` directory.

This step is important, because the conversion scripts reference this README.md for this info. Otherwise the books will not have the proper titles in the database or other formats. 

## Step Four: Add the zip.

The zip should be placed in the `<translation>` directory. The case of the directory must match the case of the zip you downloaded. In this case it looks like this: 

`sources/hbo/WLC/WLC.zip`

This is necessary because the conversion scripts rely on this convention to find all necessary scripts and info.

## Step Six: Run scripts/extract_esword_zips.py

Now we run `python extract_esword_zips.py` in the `scripts` directory. 

This will loop over all translations directories. If a `<translation>.json` file does not exist in a directory, it will run
the `sword_to_json.py` script on the `<translation>.zip` file. 

> **NOTE:** The process can crash depending on how well the actual sword_to_json packages works on a given .zip. Generally it seems
to run well, but be warned. 

## Step Seven: Check the `<translation>.json` file...

Check to see that the conversion went well in a text document. If all went well, you should see the JSON object populated with text.
> **NOTE** some translations can have missing books / chapters / verses. Common cases is translations where either old testament or new testament is left out.

## Now the base text is ready for conversion by other scripts...

One of the rules we have for this project: NEVER ALTER the .json files exported from the ESWORD export. This is so that we do not 
risk changing anything wrongly. If by some minor chance you find an error in ESWORD itself (either that it itself had, or an export issue) then change the results in your database / text document. 

You can now run the other scripts to create other formats. You can find these and their functions in the documentation.
