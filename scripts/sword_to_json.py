from pysword.modules import SwordModules
import argparse, json, sys

if sys.version_info > (3, 0):
    from past.builtins import xrange


def generate_dict(source_file, bible_version):
    modules = SwordModules(source_file)
    found_modules = modules.parse_modules()
    bible = modules.get_bible_from_module(bible_version)

    books = bible.get_structure()._books['ot'] + bible.get_structure()._books['nt']

    bib = {}
    bib['books'] = []

    for book in books:
        chapters = []
        for chapter in xrange(1, book.num_chapters+1):
            verses = []
            for verse in xrange(1, len(book.get_indicies(chapter))+1 ):
                verses.append({
                    'verse': verse,
                    'chapter': chapter,
                    'name': book.name + " " + str(chapter) + ":" + str(verse),
                    'text': bible.get(books=[book.name], chapters=[chapter], verses=[verse])
                    })
            chapters.append({
                'chapter': chapter,
                'name': book.name + " " + str(chapter),
                'verses': verses
            })
        bib['books'].append({
            'name': book.name,
            'chapters': chapters
        })
    
    return bib

def write_json(bible_dict, output_file):
    with open(output_file, 'w') as outfile:  
        json.dump(bible_dict, outfile, indent=4)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--source_file')
    parser.add_argument('--bible_version')
    parser.add_argument('--output_file')
    args = parser.parse_args()

    bible_dict = generate_dict(args.source_file, args.bible_version)
    write_json(bible_dict, args.output_file)

if __name__ == "__main__": main()
