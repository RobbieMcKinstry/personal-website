# Graduate School

It's no secret I'm planning to pursue a graduate education. My research interests are in
program synthesis, automated defect repair, program analysis, type theory and language theory.

If you're a principal investigator working in the above, and you're looking for a new 
graduate student, feel free to get in contact with me.

# Current Research: HashiCorp

I'm a full-time research engineer at HashiCorp. At HashiCorp, I work on presently unannounced security
research. We'll likely annouce the project before the end of the year, and publish sometime after that.
We're also in the processing of submitting a patent.

I began at HashiCorp on 15 May 2017 after my graduation from the University of Pittsburgh. I
joined in the first 100 concurrent employees, and as the first research engineer. There are two 
permanent members of the research team: 
[Jon Currey](https://scholar.google.com/citations?user=i6jP7qMAAAAJ) and myself. We are overseen by
[Armon Dadgar](https://twitter.com/armon), HashiCorp's co-founder.

Here's proof: 

<blockquote class="twitter-tweet" data-lang="en" align="center"><p lang="en" dir="ltr">HashiCorp Research sync. Fun to get the distributed team together and see progress and discuss future work. <a href="https://t.co/tfPsR55Muq">pic.twitter.com/tfPsR55Muq</a></p>&mdash; Mitchell Hashimoto (@mitchellh) <a href="https://twitter.com/mitchellh/status/917830848170414081?ref_src=twsrc%5Etfw">October 10, 2017</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

I'm on the left in the pink. This was at the yearly off-site for the Office of the CTO.

HashiCorp's remote-first culture allows me to teach first thing in the morning, a vocation
I greatly enjoy. I've always loved the HashiCorp products since my first internship at ModCloth.
It's an incredible pleasure to work so closely with such excellent engineers, and to build my
research portfolio as well.

# Prior Research

## Wyvern

Previously, I worked in the lab of [Jonathan Aldrich](http://www.cs.cmu.edu/~aldrich/) at 
Carnegie-Mellon for two years. Dr. Aldrich is most well-known for his work on
[typestate](https://en.wikipedia.org/wiki/Typestate_analysis) type-systems for Object 
Oriented Languages (see [Plaid](https://www.cs.cmu.edu/~aldrich/plaid/).

Under Dr. Aldrich and with the mentorship of 
[Dr. Alex Potanin](http://homepages.mcs.vuw.ac.nz/~alex/), I worked on the
[Wyvern Programming Language](https://wyvernlang.github.io/). Wyvern is a implicitly-typed language,
which means it's strongly typed but doesn't have classes. Instead, when an object is created, its
members define its type. A object may only be stored in a variable of type `Foo` if that object
contains all of the methods and fields defined in `Foo`. 

I'd love to see more language pursue implicit typing. I feel like implicit typing
has all of the readability benefits of dynamic languages while all of the correctness
benefits of static languages. Ironically, in my work with Alex and Jonathan, Wyvern's
implicit typing was never emphasized. They seemed to feel that the real strengths of Wyvern
lay elsewhere.

## Omega

My first exposure to computer science research was under Dr. David Fisher, one of the minds
behind Ada. Under Dr. Fisher, I implemented an Earley Parser in C for use in the Omega programming 
language. Omega has a property-based type system and its designed to reflect natural language. 
Earley Parsers like like other context-free grammar parses, except that they can parse any
context-free grammar. Most textbook-taught parsing algorithms have restrictions on the grammar,
like that it can't be left-recurvie, or that it can't have multiple valid parses.
Earley's Algorithm allows a parser to parse sentences that have infinite valid parses by representing
the parse tree as a "parse graph", where cycles indicate ambiguity. Additionally, Earley's Algorithm
is also highly performant, utilizing dynamic programming to cache parse matches.

I highly recommend [this article](https://jeffreykegler.github.io/personal/timeline_v3) on parsing 
by Jeffrey Kegler if you want to know more.
