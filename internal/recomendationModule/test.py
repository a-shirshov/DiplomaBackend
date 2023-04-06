import spacy
import numpy as np

nlp = spacy.load("ru_core_news_md")

doc1 = nlp("Уличные коты опять подрались с собаками на улице")

doc2 = nlp("Кошки никак не могут ужиться с собаками")

print(doc1.similarity(doc2))

lemmatized_text1 = ' '.join([token.lemma_ for token in doc1])
print(lemmatized_text1)

lemmatized_text2 = ' '.join([token.lemma_ for token in doc2])
print(lemmatized_text2)

doc3 = nlp(lemmatized_text1)
doc4 = nlp(lemmatized_text2)

print(doc3.similarity(doc4))

lemmatized_text1 = ' '.join([token.lemma_ for token in doc1 if not token.is_stop])
print(lemmatized_text1)

lemmatized_text2 = ' '.join([token.lemma_ for token in doc2 if not token.is_stop])
print(lemmatized_text2)

doc3 = nlp(lemmatized_text1)
doc4 = nlp(lemmatized_text2)

print(doc3.similarity(doc4))