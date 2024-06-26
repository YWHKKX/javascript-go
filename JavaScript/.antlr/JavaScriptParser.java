// Generated from /home/yhellow/桌面/javascript-go/src/JavaScript/JavaScriptParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class JavaScriptParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		FUNCTION=1, RETURN=2, VAR=3, CLASS=4, THIS=5, CONS=6, NEW=7, NULL=8, UNDEFINED=9, 
		IF=10, ELSE=11, WHILE=12, FOR=13, BREAK=14, CONTINUE=15, CASE=16, SWITCH=17, 
		DEFAULT=18, TRUE=19, FALSE=20, LPAREN=21, RPAREN=22, LBRACE=23, RBRACE=24, 
		LBRACKET=25, RBRACKET=26, NUMBER=27, IDENTIFIER=28, STRING=29, COL=30, 
		DOT=31, COMMA=32, SEMICOLON=33, ASSIGN=34, ADD=35, SUB=36, MUL=37, DIV=38, 
		MOD=39, INC=40, DEC=41, ADD_ASSIGN=42, SUB_ASSIGN=43, MUL_ASSIGN=44, DIV_ASSIGN=45, 
		MOD_ASSIGN=46, NOT=47, EQ=48, NEQ=49, LT=50, GT=51, LTE=52, GTE=53, AEQ=54, 
		NAEQ=55, AND=56, OR=57, WS=58, COMMENT=59, LINE_COMMENT=60;
	public static final int
		RULE_num = 0, RULE_id = 1, RULE_str = 2, RULE_arr = 3, RULE_key = 4, RULE_obj = 5, 
		RULE_bool = 6, RULE_this = 7, RULE_funcdef = 8, RULE_paramlist = 9, RULE_param = 10, 
		RULE_funcall = 11, RULE_exprlist = 12, RULE_cobj = 13, RULE_class = 14, 
		RULE_cons = 15, RULE_method = 16, RULE_program = 17, RULE_global = 18, 
		RULE_expr = 19, RULE_stmg = 20, RULE_stm = 21, RULE_case = 22, RULE_default = 23, 
		RULE_block = 24;
	private static String[] makeRuleNames() {
		return new String[] {
			"num", "id", "str", "arr", "key", "obj", "bool", "this", "funcdef", "paramlist", 
			"param", "funcall", "exprlist", "cobj", "class", "cons", "method", "program", 
			"global", "expr", "stmg", "stm", "case", "default", "block"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'function'", "'return'", "'var'", "'class'", "'this'", "'constructor'", 
			"'new'", "'null'", "'undefined'", "'if'", "'else'", "'while'", "'for'", 
			"'break'", "'continue'", "'case'", "'switch'", "'default'", "'true'", 
			"'false'", "'('", "')'", "'{'", "'}'", "'['", "']'", null, null, null, 
			"':'", "'.'", "','", "';'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", 
			"'++'", "'--'", "'+='", "'-='", "'*='", "'/='", "'%='", "'!'", "'=='", 
			"'!='", "'<'", "'>'", "'<='", "'>='", "'==='", "'!=='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "FUNCTION", "RETURN", "VAR", "CLASS", "THIS", "CONS", "NEW", "NULL", 
			"UNDEFINED", "IF", "ELSE", "WHILE", "FOR", "BREAK", "CONTINUE", "CASE", 
			"SWITCH", "DEFAULT", "TRUE", "FALSE", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
			"LBRACKET", "RBRACKET", "NUMBER", "IDENTIFIER", "STRING", "COL", "DOT", 
			"COMMA", "SEMICOLON", "ASSIGN", "ADD", "SUB", "MUL", "DIV", "MOD", "INC", 
			"DEC", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", 
			"NOT", "EQ", "NEQ", "LT", "GT", "LTE", "GTE", "AEQ", "NAEQ", "AND", "OR", 
			"WS", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "JavaScriptParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public JavaScriptParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NumContext extends ParserRuleContext {
		public TerminalNode NUMBER() { return getToken(JavaScriptParser.NUMBER, 0); }
		public NumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_num; }
	}

	public final NumContext num() throws RecognitionException {
		NumContext _localctx = new NumContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_num);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			match(NUMBER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public IdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_id; }
	}

	public final IdContext id() throws RecognitionException {
		IdContext _localctx = new IdContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_id);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StrContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(JavaScriptParser.STRING, 0); }
		public StrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_str; }
	}

	public final StrContext str() throws RecognitionException {
		StrContext _localctx = new StrContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_str);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrContext extends ParserRuleContext {
		public TerminalNode LBRACKET() { return getToken(JavaScriptParser.LBRACKET, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode RBRACKET() { return getToken(JavaScriptParser.RBRACKET, 0); }
		public List<TerminalNode> COMMA() { return getTokens(JavaScriptParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(JavaScriptParser.COMMA, i);
		}
		public ArrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arr; }
	}

	public final ArrContext arr() throws RecognitionException {
		ArrContext _localctx = new ArrContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_arr);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			match(LBRACKET);
			setState(57);
			expr(0);
			setState(62);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(58);
				match(COMMA);
				setState(59);
				expr(0);
				}
				}
				setState(64);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(65);
			match(RBRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeyContext extends ParserRuleContext {
		public IdContext id() {
			return getRuleContext(IdContext.class,0);
		}
		public TerminalNode COL() { return getToken(JavaScriptParser.COL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(67);
			id();
			setState(68);
			match(COL);
			setState(69);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ObjContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(JavaScriptParser.LBRACE, 0); }
		public List<KeyContext> key() {
			return getRuleContexts(KeyContext.class);
		}
		public KeyContext key(int i) {
			return getRuleContext(KeyContext.class,i);
		}
		public TerminalNode RBRACE() { return getToken(JavaScriptParser.RBRACE, 0); }
		public List<TerminalNode> COMMA() { return getTokens(JavaScriptParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(JavaScriptParser.COMMA, i);
		}
		public ObjContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_obj; }
	}

	public final ObjContext obj() throws RecognitionException {
		ObjContext _localctx = new ObjContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_obj);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(LBRACE);
			setState(72);
			key();
			setState(77);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(73);
				match(COMMA);
				setState(74);
				key();
				}
				}
				setState(79);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(80);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BoolContext extends ParserRuleContext {
		public TerminalNode TRUE() { return getToken(JavaScriptParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(JavaScriptParser.FALSE, 0); }
		public BoolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool; }
	}

	public final BoolContext bool() throws RecognitionException {
		BoolContext _localctx = new BoolContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_bool);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			_la = _input.LA(1);
			if ( !(_la==TRUE || _la==FALSE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ThisContext extends ParserRuleContext {
		public TerminalNode THIS() { return getToken(JavaScriptParser.THIS, 0); }
		public ThisContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_this; }
	}

	public final ThisContext this_() throws RecognitionException {
		ThisContext _localctx = new ThisContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_this);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			match(THIS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncdefContext extends ParserRuleContext {
		public TerminalNode FUNCTION() { return getToken(JavaScriptParser.FUNCTION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ParamlistContext paramlist() {
			return getRuleContext(ParamlistContext.class,0);
		}
		public FuncdefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcdef; }
	}

	public final FuncdefContext funcdef() throws RecognitionException {
		FuncdefContext _localctx = new FuncdefContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_funcdef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			match(FUNCTION);
			setState(87);
			match(IDENTIFIER);
			setState(88);
			match(LPAREN);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(89);
				paramlist();
				}
			}

			setState(92);
			match(RPAREN);
			setState(93);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamlistContext extends ParserRuleContext {
		public List<ParamContext> param() {
			return getRuleContexts(ParamContext.class);
		}
		public ParamContext param(int i) {
			return getRuleContext(ParamContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(JavaScriptParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(JavaScriptParser.COMMA, i);
		}
		public ParamlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramlist; }
	}

	public final ParamlistContext paramlist() throws RecognitionException {
		ParamlistContext _localctx = new ParamlistContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_paramlist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			param();
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(96);
				match(COMMA);
				setState(97);
				param();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParamContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode ASSIGN() { return getToken(JavaScriptParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_param);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			match(IDENTIFIER);
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(104);
				match(ASSIGN);
				setState(105);
				expr(0);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncallContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public ExprlistContext exprlist() {
			return getRuleContext(ExprlistContext.class,0);
		}
		public FuncallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcall; }
	}

	public final FuncallContext funcall() throws RecognitionException {
		FuncallContext _localctx = new FuncallContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_funcall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(IDENTIFIER);
			setState(109);
			match(LPAREN);
			setState(111);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726279840L) != 0)) {
				{
				setState(110);
				exprlist();
				}
			}

			setState(113);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprlistContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(JavaScriptParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(JavaScriptParser.COMMA, i);
		}
		public ExprlistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprlist; }
	}

	public final ExprlistContext exprlist() throws RecognitionException {
		ExprlistContext _localctx = new ExprlistContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_exprlist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			expr(0);
			setState(120);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(116);
				match(COMMA);
				setState(117);
				expr(0);
				}
				}
				setState(122);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CobjContext extends ParserRuleContext {
		public TerminalNode NEW() { return getToken(JavaScriptParser.NEW, 0); }
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public ExprlistContext exprlist() {
			return getRuleContext(ExprlistContext.class,0);
		}
		public CobjContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cobj; }
	}

	public final CobjContext cobj() throws RecognitionException {
		CobjContext _localctx = new CobjContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_cobj);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(NEW);
			setState(124);
			match(IDENTIFIER);
			setState(125);
			match(LPAREN);
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726279840L) != 0)) {
				{
				setState(126);
				exprlist();
				}
			}

			setState(129);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ClassContext extends ParserRuleContext {
		public TerminalNode CLASS() { return getToken(JavaScriptParser.CLASS, 0); }
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode LBRACE() { return getToken(JavaScriptParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(JavaScriptParser.RBRACE, 0); }
		public ConsContext cons() {
			return getRuleContext(ConsContext.class,0);
		}
		public List<MethodContext> method() {
			return getRuleContexts(MethodContext.class);
		}
		public MethodContext method(int i) {
			return getRuleContext(MethodContext.class,i);
		}
		public ClassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_class; }
	}

	public final ClassContext class_() throws RecognitionException {
		ClassContext _localctx = new ClassContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_class);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			match(CLASS);
			setState(132);
			match(IDENTIFIER);
			setState(133);
			match(LBRACE);
			setState(135);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==CONS) {
				{
				setState(134);
				cons();
				}
			}

			setState(140);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(137);
				method();
				}
				}
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(143);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ConsContext extends ParserRuleContext {
		public TerminalNode CONS() { return getToken(JavaScriptParser.CONS, 0); }
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ParamlistContext paramlist() {
			return getRuleContext(ParamlistContext.class,0);
		}
		public ConsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cons; }
	}

	public final ConsContext cons() throws RecognitionException {
		ConsContext _localctx = new ConsContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_cons);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(145);
			match(CONS);
			setState(146);
			match(LPAREN);
			setState(148);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(147);
				paramlist();
				}
			}

			setState(150);
			match(RPAREN);
			setState(151);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ParamlistContext paramlist() {
			return getRuleContext(ParamlistContext.class,0);
		}
		public MethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method; }
	}

	public final MethodContext method() throws RecognitionException {
		MethodContext _localctx = new MethodContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_method);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(153);
			match(IDENTIFIER);
			setState(154);
			match(LPAREN);
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(155);
				paramlist();
				}
			}

			setState(158);
			match(RPAREN);
			setState(159);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<GlobalContext> global() {
			return getRuleContexts(GlobalContext.class);
		}
		public GlobalContext global(int i) {
			return getRuleContext(GlobalContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473406L) != 0)) {
				{
				{
				setState(161);
				global();
				}
				}
				setState(166);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GlobalContext extends ParserRuleContext {
		public FuncdefContext funcdef() {
			return getRuleContext(FuncdefContext.class,0);
		}
		public StmgContext stmg() {
			return getRuleContext(StmgContext.class,0);
		}
		public ClassContext class_() {
			return getRuleContext(ClassContext.class,0);
		}
		public GlobalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_global; }
	}

	public final GlobalContext global() throws RecognitionException {
		GlobalContext _localctx = new GlobalContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_global);
		try {
			setState(170);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNCTION:
				enterOuterAlt(_localctx, 1);
				{
				setState(167);
				funcdef();
				}
				break;
			case RETURN:
			case VAR:
			case THIS:
			case NEW:
			case IF:
			case WHILE:
			case FOR:
			case BREAK:
			case CONTINUE:
			case SWITCH:
			case LPAREN:
			case LBRACE:
			case LBRACKET:
			case NUMBER:
			case IDENTIFIER:
			case STRING:
			case SUB:
			case INC:
			case DEC:
			case NOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(168);
				stmg();
				}
				break;
			case CLASS:
				enterOuterAlt(_localctx, 3);
				{
				setState(169);
				class_();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public TerminalNode INC() { return getToken(JavaScriptParser.INC, 0); }
		public TerminalNode DEC() { return getToken(JavaScriptParser.DEC, 0); }
		public TerminalNode NOT() { return getToken(JavaScriptParser.NOT, 0); }
		public TerminalNode SUB() { return getToken(JavaScriptParser.SUB, 0); }
		public ThisContext this_() {
			return getRuleContext(ThisContext.class,0);
		}
		public CobjContext cobj() {
			return getRuleContext(CobjContext.class,0);
		}
		public FuncallContext funcall() {
			return getRuleContext(FuncallContext.class,0);
		}
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public IdContext id() {
			return getRuleContext(IdContext.class,0);
		}
		public StrContext str() {
			return getRuleContext(StrContext.class,0);
		}
		public ArrContext arr() {
			return getRuleContext(ArrContext.class,0);
		}
		public ObjContext obj() {
			return getRuleContext(ObjContext.class,0);
		}
		public TerminalNode DOT() { return getToken(JavaScriptParser.DOT, 0); }
		public TerminalNode MUL() { return getToken(JavaScriptParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(JavaScriptParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(JavaScriptParser.ADD, 0); }
		public TerminalNode LT() { return getToken(JavaScriptParser.LT, 0); }
		public TerminalNode GT() { return getToken(JavaScriptParser.GT, 0); }
		public TerminalNode EQ() { return getToken(JavaScriptParser.EQ, 0); }
		public TerminalNode LTE() { return getToken(JavaScriptParser.LTE, 0); }
		public TerminalNode GTE() { return getToken(JavaScriptParser.GTE, 0); }
		public TerminalNode NEQ() { return getToken(JavaScriptParser.NEQ, 0); }
		public TerminalNode AEQ() { return getToken(JavaScriptParser.AEQ, 0); }
		public TerminalNode NAEQ() { return getToken(JavaScriptParser.NAEQ, 0); }
		public TerminalNode AND() { return getToken(JavaScriptParser.AND, 0); }
		public TerminalNode OR() { return getToken(JavaScriptParser.OR, 0); }
		public TerminalNode LBRACKET() { return getToken(JavaScriptParser.LBRACKET, 0); }
		public TerminalNode RBRACKET() { return getToken(JavaScriptParser.RBRACKET, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(173);
				match(LPAREN);
				setState(174);
				expr(0);
				setState(175);
				match(RPAREN);
				}
				break;
			case 2:
				{
				setState(177);
				_la = _input.LA(1);
				if ( !(_la==INC || _la==DEC) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(178);
				expr(11);
				}
				break;
			case 3:
				{
				setState(179);
				_la = _input.LA(1);
				if ( !(_la==SUB || _la==NOT) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(180);
				expr(9);
				}
				break;
			case 4:
				{
				setState(181);
				this_();
				}
				break;
			case 5:
				{
				setState(182);
				cobj();
				}
				break;
			case 6:
				{
				setState(183);
				funcall();
				}
				break;
			case 7:
				{
				setState(184);
				num();
				}
				break;
			case 8:
				{
				setState(185);
				id();
				}
				break;
			case 9:
				{
				setState(186);
				str();
				}
				break;
			case 10:
				{
				setState(187);
				arr();
				}
				break;
			case 11:
				{
				setState(188);
				obj();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(212);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(210);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(191);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(192);
						match(DOT);
						setState(193);
						expr(17);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(194);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(195);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 515396075520L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(196);
						expr(15);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(197);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(198);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 71776119061217280L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(199);
						expr(14);
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(200);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(201);
						_la = _input.LA(1);
						if ( !(_la==AND || _la==OR) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(202);
						expr(13);
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(203);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(204);
						match(LBRACKET);
						setState(205);
						expr(0);
						setState(206);
						match(RBRACKET);
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(208);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(209);
						_la = _input.LA(1);
						if ( !(_la==INC || _la==DEC) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						}
						break;
					}
					} 
				}
				setState(214);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmgContext extends ParserRuleContext {
		public StmContext stm() {
			return getRuleContext(StmContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(JavaScriptParser.SEMICOLON, 0); }
		public StmgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmg; }
	}

	public final StmgContext stmg() throws RecognitionException {
		StmgContext _localctx = new StmgContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_stmg);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			stm();
			setState(217);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(216);
				match(SEMICOLON);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(JavaScriptParser.ASSIGN, 0); }
		public TerminalNode ADD_ASSIGN() { return getToken(JavaScriptParser.ADD_ASSIGN, 0); }
		public TerminalNode SUB_ASSIGN() { return getToken(JavaScriptParser.SUB_ASSIGN, 0); }
		public TerminalNode MUL_ASSIGN() { return getToken(JavaScriptParser.MUL_ASSIGN, 0); }
		public TerminalNode DIV_ASSIGN() { return getToken(JavaScriptParser.DIV_ASSIGN, 0); }
		public TerminalNode VAR() { return getToken(JavaScriptParser.VAR, 0); }
		public TerminalNode IDENTIFIER() { return getToken(JavaScriptParser.IDENTIFIER, 0); }
		public TerminalNode RETURN() { return getToken(JavaScriptParser.RETURN, 0); }
		public TerminalNode BREAK() { return getToken(JavaScriptParser.BREAK, 0); }
		public TerminalNode CONTINUE() { return getToken(JavaScriptParser.CONTINUE, 0); }
		public TerminalNode IF() { return getToken(JavaScriptParser.IF, 0); }
		public TerminalNode LPAREN() { return getToken(JavaScriptParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(JavaScriptParser.RPAREN, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(JavaScriptParser.ELSE, 0); }
		public TerminalNode WHILE() { return getToken(JavaScriptParser.WHILE, 0); }
		public TerminalNode FOR() { return getToken(JavaScriptParser.FOR, 0); }
		public List<TerminalNode> SEMICOLON() { return getTokens(JavaScriptParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(JavaScriptParser.SEMICOLON, i);
		}
		public List<StmContext> stm() {
			return getRuleContexts(StmContext.class);
		}
		public StmContext stm(int i) {
			return getRuleContext(StmContext.class,i);
		}
		public TerminalNode SWITCH() { return getToken(JavaScriptParser.SWITCH, 0); }
		public TerminalNode LBRACE() { return getToken(JavaScriptParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(JavaScriptParser.RBRACE, 0); }
		public List<CaseContext> case_() {
			return getRuleContexts(CaseContext.class);
		}
		public CaseContext case_(int i) {
			return getRuleContext(CaseContext.class,i);
		}
		public DefaultContext default_() {
			return getRuleContext(DefaultContext.class,0);
		}
		public StmContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stm; }
	}

	public final StmContext stm() throws RecognitionException {
		StmContext _localctx = new StmContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_stm);
		int _la;
		try {
			setState(281);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(219);
				expr(0);
				setState(220);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 65987877535744L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(221);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(223);
				match(VAR);
				setState(224);
				match(IDENTIFIER);
				setState(227);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(225);
					match(ASSIGN);
					setState(226);
					expr(0);
					}
				}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(229);
				match(RETURN);
				setState(230);
				expr(0);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(231);
				match(BREAK);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(232);
				match(CONTINUE);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(233);
				match(IF);
				setState(234);
				match(LPAREN);
				setState(235);
				expr(0);
				setState(236);
				match(RPAREN);
				setState(237);
				block();
				setState(240);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ELSE) {
					{
					setState(238);
					match(ELSE);
					setState(239);
					block();
					}
				}

				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(242);
				match(WHILE);
				setState(243);
				match(LPAREN);
				setState(244);
				expr(0);
				setState(245);
				match(RPAREN);
				setState(246);
				block();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(248);
				match(FOR);
				setState(249);
				match(LPAREN);
				setState(251);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473388L) != 0)) {
					{
					setState(250);
					stm();
					}
				}

				setState(253);
				match(SEMICOLON);
				setState(255);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473388L) != 0)) {
					{
					setState(254);
					stm();
					}
				}

				setState(257);
				match(SEMICOLON);
				setState(259);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473388L) != 0)) {
					{
					setState(258);
					stm();
					}
				}

				setState(261);
				match(RPAREN);
				setState(262);
				block();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(263);
				match(SWITCH);
				setState(264);
				match(LPAREN);
				setState(265);
				expr(0);
				setState(266);
				match(RPAREN);
				setState(267);
				match(LBRACE);
				setState(271);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==CASE) {
					{
					{
					setState(268);
					case_();
					}
					}
					setState(273);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(275);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==DEFAULT) {
					{
					setState(274);
					default_();
					}
				}

				setState(277);
				match(RBRACE);
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(279);
				block();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(280);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CaseContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(JavaScriptParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COL() { return getToken(JavaScriptParser.COL, 0); }
		public List<StmContext> stm() {
			return getRuleContexts(StmContext.class);
		}
		public StmContext stm(int i) {
			return getRuleContext(StmContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(JavaScriptParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(JavaScriptParser.SEMICOLON, i);
		}
		public CaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case; }
	}

	public final CaseContext case_() throws RecognitionException {
		CaseContext _localctx = new CaseContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			match(CASE);
			setState(284);
			expr(0);
			setState(285);
			match(COL);
			setState(292);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473388L) != 0)) {
				{
				{
				setState(286);
				stm();
				setState(288);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(287);
					match(SEMICOLON);
					}
				}

				}
				}
				setState(294);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefaultContext extends ParserRuleContext {
		public TerminalNode DEFAULT() { return getToken(JavaScriptParser.DEFAULT, 0); }
		public TerminalNode COL() { return getToken(JavaScriptParser.COL, 0); }
		public List<StmContext> stm() {
			return getRuleContexts(StmContext.class);
		}
		public StmContext stm(int i) {
			return getRuleContext(StmContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(JavaScriptParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(JavaScriptParser.SEMICOLON, i);
		}
		public DefaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default; }
	}

	public final DefaultContext default_() throws RecognitionException {
		DefaultContext _localctx = new DefaultContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_default);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			match(DEFAULT);
			setState(296);
			match(COL);
			setState(303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473388L) != 0)) {
				{
				{
				setState(297);
				stm();
				setState(299);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(298);
					match(SEMICOLON);
					}
				}

				}
				}
				setState(305);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(JavaScriptParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(JavaScriptParser.RBRACE, 0); }
		public List<StmContext> stm() {
			return getRuleContexts(StmContext.class);
		}
		public StmContext stm(int i) {
			return getRuleContext(StmContext.class,i);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(JavaScriptParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(JavaScriptParser.SEMICOLON, i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			match(LBRACE);
			setState(313);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144105726473388L) != 0)) {
				{
				{
				setState(307);
				stm();
				setState(309);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(308);
					match(SEMICOLON);
					}
				}

				}
				}
				setState(315);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(316);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 19:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 16);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 15);
		case 5:
			return precpred(_ctx, 10);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001<\u013f\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003=\b\u0003"+
		"\n\u0003\f\u0003@\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0005\u0005L\b\u0005\n\u0005\f\u0005O\t\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\b[\b\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001"+
		"\t\u0005\tc\b\t\n\t\f\tf\t\t\u0001\n\u0001\n\u0001\n\u0003\nk\b\n\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0003\u000bp\b\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\f\u0001\f\u0001\f\u0005\fw\b\f\n\f\f\fz\t\f\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0003\r\u0080\b\r\u0001\r\u0001\r\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u0088\b\u000e\u0001\u000e\u0005"+
		"\u000e\u008b\b\u000e\n\u000e\f\u000e\u008e\t\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0095\b\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010"+
		"\u009d\b\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0005\u0011"+
		"\u00a3\b\u0011\n\u0011\f\u0011\u00a6\t\u0011\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0003\u0012\u00ab\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0003\u0013\u00be\b\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u00d3"+
		"\b\u0013\n\u0013\f\u0013\u00d6\t\u0013\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u00da\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u00e4\b\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u00f1\b\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u00fc\b\u0015\u0001\u0015"+
		"\u0001\u0015\u0003\u0015\u0100\b\u0015\u0001\u0015\u0001\u0015\u0003\u0015"+
		"\u0104\b\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u010e\b\u0015\n\u0015"+
		"\f\u0015\u0111\t\u0015\u0001\u0015\u0003\u0015\u0114\b\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u011a\b\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u0121\b\u0016"+
		"\u0005\u0016\u0123\b\u0016\n\u0016\f\u0016\u0126\t\u0016\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u012c\b\u0017\u0005\u0017\u012e"+
		"\b\u0017\n\u0017\f\u0017\u0131\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0003\u0018\u0136\b\u0018\u0005\u0018\u0138\b\u0018\n\u0018\f\u0018\u013b"+
		"\t\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0000\u0001&\u0019\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.0\u0000\u0007\u0001\u0000\u0013\u0014\u0001\u0000()\u0002"+
		"\u0000$$//\u0001\u0000#&\u0001\u000007\u0001\u000089\u0002\u0000\"\"*"+
		"-\u015c\u00002\u0001\u0000\u0000\u0000\u00024\u0001\u0000\u0000\u0000"+
		"\u00046\u0001\u0000\u0000\u0000\u00068\u0001\u0000\u0000\u0000\bC\u0001"+
		"\u0000\u0000\u0000\nG\u0001\u0000\u0000\u0000\fR\u0001\u0000\u0000\u0000"+
		"\u000eT\u0001\u0000\u0000\u0000\u0010V\u0001\u0000\u0000\u0000\u0012_"+
		"\u0001\u0000\u0000\u0000\u0014g\u0001\u0000\u0000\u0000\u0016l\u0001\u0000"+
		"\u0000\u0000\u0018s\u0001\u0000\u0000\u0000\u001a{\u0001\u0000\u0000\u0000"+
		"\u001c\u0083\u0001\u0000\u0000\u0000\u001e\u0091\u0001\u0000\u0000\u0000"+
		" \u0099\u0001\u0000\u0000\u0000\"\u00a4\u0001\u0000\u0000\u0000$\u00aa"+
		"\u0001\u0000\u0000\u0000&\u00bd\u0001\u0000\u0000\u0000(\u00d7\u0001\u0000"+
		"\u0000\u0000*\u0119\u0001\u0000\u0000\u0000,\u011b\u0001\u0000\u0000\u0000"+
		".\u0127\u0001\u0000\u0000\u00000\u0132\u0001\u0000\u0000\u000023\u0005"+
		"\u001b\u0000\u00003\u0001\u0001\u0000\u0000\u000045\u0005\u001c\u0000"+
		"\u00005\u0003\u0001\u0000\u0000\u000067\u0005\u001d\u0000\u00007\u0005"+
		"\u0001\u0000\u0000\u000089\u0005\u0019\u0000\u00009>\u0003&\u0013\u0000"+
		":;\u0005 \u0000\u0000;=\u0003&\u0013\u0000<:\u0001\u0000\u0000\u0000="+
		"@\u0001\u0000\u0000\u0000><\u0001\u0000\u0000\u0000>?\u0001\u0000\u0000"+
		"\u0000?A\u0001\u0000\u0000\u0000@>\u0001\u0000\u0000\u0000AB\u0005\u001a"+
		"\u0000\u0000B\u0007\u0001\u0000\u0000\u0000CD\u0003\u0002\u0001\u0000"+
		"DE\u0005\u001e\u0000\u0000EF\u0003&\u0013\u0000F\t\u0001\u0000\u0000\u0000"+
		"GH\u0005\u0017\u0000\u0000HM\u0003\b\u0004\u0000IJ\u0005 \u0000\u0000"+
		"JL\u0003\b\u0004\u0000KI\u0001\u0000\u0000\u0000LO\u0001\u0000\u0000\u0000"+
		"MK\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000\u0000NP\u0001\u0000\u0000"+
		"\u0000OM\u0001\u0000\u0000\u0000PQ\u0005\u0018\u0000\u0000Q\u000b\u0001"+
		"\u0000\u0000\u0000RS\u0007\u0000\u0000\u0000S\r\u0001\u0000\u0000\u0000"+
		"TU\u0005\u0005\u0000\u0000U\u000f\u0001\u0000\u0000\u0000VW\u0005\u0001"+
		"\u0000\u0000WX\u0005\u001c\u0000\u0000XZ\u0005\u0015\u0000\u0000Y[\u0003"+
		"\u0012\t\u0000ZY\u0001\u0000\u0000\u0000Z[\u0001\u0000\u0000\u0000[\\"+
		"\u0001\u0000\u0000\u0000\\]\u0005\u0016\u0000\u0000]^\u00030\u0018\u0000"+
		"^\u0011\u0001\u0000\u0000\u0000_d\u0003\u0014\n\u0000`a\u0005 \u0000\u0000"+
		"ac\u0003\u0014\n\u0000b`\u0001\u0000\u0000\u0000cf\u0001\u0000\u0000\u0000"+
		"db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000e\u0013\u0001\u0000"+
		"\u0000\u0000fd\u0001\u0000\u0000\u0000gj\u0005\u001c\u0000\u0000hi\u0005"+
		"\"\u0000\u0000ik\u0003&\u0013\u0000jh\u0001\u0000\u0000\u0000jk\u0001"+
		"\u0000\u0000\u0000k\u0015\u0001\u0000\u0000\u0000lm\u0005\u001c\u0000"+
		"\u0000mo\u0005\u0015\u0000\u0000np\u0003\u0018\f\u0000on\u0001\u0000\u0000"+
		"\u0000op\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qr\u0005\u0016"+
		"\u0000\u0000r\u0017\u0001\u0000\u0000\u0000sx\u0003&\u0013\u0000tu\u0005"+
		" \u0000\u0000uw\u0003&\u0013\u0000vt\u0001\u0000\u0000\u0000wz\u0001\u0000"+
		"\u0000\u0000xv\u0001\u0000\u0000\u0000xy\u0001\u0000\u0000\u0000y\u0019"+
		"\u0001\u0000\u0000\u0000zx\u0001\u0000\u0000\u0000{|\u0005\u0007\u0000"+
		"\u0000|}\u0005\u001c\u0000\u0000}\u007f\u0005\u0015\u0000\u0000~\u0080"+
		"\u0003\u0018\f\u0000\u007f~\u0001\u0000\u0000\u0000\u007f\u0080\u0001"+
		"\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000\u0000\u0081\u0082\u0005"+
		"\u0016\u0000\u0000\u0082\u001b\u0001\u0000\u0000\u0000\u0083\u0084\u0005"+
		"\u0004\u0000\u0000\u0084\u0085\u0005\u001c\u0000\u0000\u0085\u0087\u0005"+
		"\u0017\u0000\u0000\u0086\u0088\u0003\u001e\u000f\u0000\u0087\u0086\u0001"+
		"\u0000\u0000\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\u008c\u0001"+
		"\u0000\u0000\u0000\u0089\u008b\u0003 \u0010\u0000\u008a\u0089\u0001\u0000"+
		"\u0000\u0000\u008b\u008e\u0001\u0000\u0000\u0000\u008c\u008a\u0001\u0000"+
		"\u0000\u0000\u008c\u008d\u0001\u0000\u0000\u0000\u008d\u008f\u0001\u0000"+
		"\u0000\u0000\u008e\u008c\u0001\u0000\u0000\u0000\u008f\u0090\u0005\u0018"+
		"\u0000\u0000\u0090\u001d\u0001\u0000\u0000\u0000\u0091\u0092\u0005\u0006"+
		"\u0000\u0000\u0092\u0094\u0005\u0015\u0000\u0000\u0093\u0095\u0003\u0012"+
		"\t\u0000\u0094\u0093\u0001\u0000\u0000\u0000\u0094\u0095\u0001\u0000\u0000"+
		"\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096\u0097\u0005\u0016\u0000"+
		"\u0000\u0097\u0098\u00030\u0018\u0000\u0098\u001f\u0001\u0000\u0000\u0000"+
		"\u0099\u009a\u0005\u001c\u0000\u0000\u009a\u009c\u0005\u0015\u0000\u0000"+
		"\u009b\u009d\u0003\u0012\t\u0000\u009c\u009b\u0001\u0000\u0000\u0000\u009c"+
		"\u009d\u0001\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000\u0000\u009e"+
		"\u009f\u0005\u0016\u0000\u0000\u009f\u00a0\u00030\u0018\u0000\u00a0!\u0001"+
		"\u0000\u0000\u0000\u00a1\u00a3\u0003$\u0012\u0000\u00a2\u00a1\u0001\u0000"+
		"\u0000\u0000\u00a3\u00a6\u0001\u0000\u0000\u0000\u00a4\u00a2\u0001\u0000"+
		"\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000\u0000\u00a5#\u0001\u0000\u0000"+
		"\u0000\u00a6\u00a4\u0001\u0000\u0000\u0000\u00a7\u00ab\u0003\u0010\b\u0000"+
		"\u00a8\u00ab\u0003(\u0014\u0000\u00a9\u00ab\u0003\u001c\u000e\u0000\u00aa"+
		"\u00a7\u0001\u0000\u0000\u0000\u00aa\u00a8\u0001\u0000\u0000\u0000\u00aa"+
		"\u00a9\u0001\u0000\u0000\u0000\u00ab%\u0001\u0000\u0000\u0000\u00ac\u00ad"+
		"\u0006\u0013\uffff\uffff\u0000\u00ad\u00ae\u0005\u0015\u0000\u0000\u00ae"+
		"\u00af\u0003&\u0013\u0000\u00af\u00b0\u0005\u0016\u0000\u0000\u00b0\u00be"+
		"\u0001\u0000\u0000\u0000\u00b1\u00b2\u0007\u0001\u0000\u0000\u00b2\u00be"+
		"\u0003&\u0013\u000b\u00b3\u00b4\u0007\u0002\u0000\u0000\u00b4\u00be\u0003"+
		"&\u0013\t\u00b5\u00be\u0003\u000e\u0007\u0000\u00b6\u00be\u0003\u001a"+
		"\r\u0000\u00b7\u00be\u0003\u0016\u000b\u0000\u00b8\u00be\u0003\u0000\u0000"+
		"\u0000\u00b9\u00be\u0003\u0002\u0001\u0000\u00ba\u00be\u0003\u0004\u0002"+
		"\u0000\u00bb\u00be\u0003\u0006\u0003\u0000\u00bc\u00be\u0003\n\u0005\u0000"+
		"\u00bd\u00ac\u0001\u0000\u0000\u0000\u00bd\u00b1\u0001\u0000\u0000\u0000"+
		"\u00bd\u00b3\u0001\u0000\u0000\u0000\u00bd\u00b5\u0001\u0000\u0000\u0000"+
		"\u00bd\u00b6\u0001\u0000\u0000\u0000\u00bd\u00b7\u0001\u0000\u0000\u0000"+
		"\u00bd\u00b8\u0001\u0000\u0000\u0000\u00bd\u00b9\u0001\u0000\u0000\u0000"+
		"\u00bd\u00ba\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000"+
		"\u00bd\u00bc\u0001\u0000\u0000\u0000\u00be\u00d4\u0001\u0000\u0000\u0000"+
		"\u00bf\u00c0\n\u0010\u0000\u0000\u00c0\u00c1\u0005\u001f\u0000\u0000\u00c1"+
		"\u00d3\u0003&\u0013\u0011\u00c2\u00c3\n\u000e\u0000\u0000\u00c3\u00c4"+
		"\u0007\u0003\u0000\u0000\u00c4\u00d3\u0003&\u0013\u000f\u00c5\u00c6\n"+
		"\r\u0000\u0000\u00c6\u00c7\u0007\u0004\u0000\u0000\u00c7\u00d3\u0003&"+
		"\u0013\u000e\u00c8\u00c9\n\f\u0000\u0000\u00c9\u00ca\u0007\u0005\u0000"+
		"\u0000\u00ca\u00d3\u0003&\u0013\r\u00cb\u00cc\n\u000f\u0000\u0000\u00cc"+
		"\u00cd\u0005\u0019\u0000\u0000\u00cd\u00ce\u0003&\u0013\u0000\u00ce\u00cf"+
		"\u0005\u001a\u0000\u0000\u00cf\u00d3\u0001\u0000\u0000\u0000\u00d0\u00d1"+
		"\n\n\u0000\u0000\u00d1\u00d3\u0007\u0001\u0000\u0000\u00d2\u00bf\u0001"+
		"\u0000\u0000\u0000\u00d2\u00c2\u0001\u0000\u0000\u0000\u00d2\u00c5\u0001"+
		"\u0000\u0000\u0000\u00d2\u00c8\u0001\u0000\u0000\u0000\u00d2\u00cb\u0001"+
		"\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d3\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d5\u0001"+
		"\u0000\u0000\u0000\u00d5\'\u0001\u0000\u0000\u0000\u00d6\u00d4\u0001\u0000"+
		"\u0000\u0000\u00d7\u00d9\u0003*\u0015\u0000\u00d8\u00da\u0005!\u0000\u0000"+
		"\u00d9\u00d8\u0001\u0000\u0000\u0000\u00d9\u00da\u0001\u0000\u0000\u0000"+
		"\u00da)\u0001\u0000\u0000\u0000\u00db\u00dc\u0003&\u0013\u0000\u00dc\u00dd"+
		"\u0007\u0006\u0000\u0000\u00dd\u00de\u0003&\u0013\u0000\u00de\u011a\u0001"+
		"\u0000\u0000\u0000\u00df\u00e0\u0005\u0003\u0000\u0000\u00e0\u00e3\u0005"+
		"\u001c\u0000\u0000\u00e1\u00e2\u0005\"\u0000\u0000\u00e2\u00e4\u0003&"+
		"\u0013\u0000\u00e3\u00e1\u0001\u0000\u0000\u0000\u00e3\u00e4\u0001\u0000"+
		"\u0000\u0000\u00e4\u011a\u0001\u0000\u0000\u0000\u00e5\u00e6\u0005\u0002"+
		"\u0000\u0000\u00e6\u011a\u0003&\u0013\u0000\u00e7\u011a\u0005\u000e\u0000"+
		"\u0000\u00e8\u011a\u0005\u000f\u0000\u0000\u00e9\u00ea\u0005\n\u0000\u0000"+
		"\u00ea\u00eb\u0005\u0015\u0000\u0000\u00eb\u00ec\u0003&\u0013\u0000\u00ec"+
		"\u00ed\u0005\u0016\u0000\u0000\u00ed\u00f0\u00030\u0018\u0000\u00ee\u00ef"+
		"\u0005\u000b\u0000\u0000\u00ef\u00f1\u00030\u0018\u0000\u00f0\u00ee\u0001"+
		"\u0000\u0000\u0000\u00f0\u00f1\u0001\u0000\u0000\u0000\u00f1\u011a\u0001"+
		"\u0000\u0000\u0000\u00f2\u00f3\u0005\f\u0000\u0000\u00f3\u00f4\u0005\u0015"+
		"\u0000\u0000\u00f4\u00f5\u0003&\u0013\u0000\u00f5\u00f6\u0005\u0016\u0000"+
		"\u0000\u00f6\u00f7\u00030\u0018\u0000\u00f7\u011a\u0001\u0000\u0000\u0000"+
		"\u00f8\u00f9\u0005\r\u0000\u0000\u00f9\u00fb\u0005\u0015\u0000\u0000\u00fa"+
		"\u00fc\u0003*\u0015\u0000\u00fb\u00fa\u0001\u0000\u0000\u0000\u00fb\u00fc"+
		"\u0001\u0000\u0000\u0000\u00fc\u00fd\u0001\u0000\u0000\u0000\u00fd\u00ff"+
		"\u0005!\u0000\u0000\u00fe\u0100\u0003*\u0015\u0000\u00ff\u00fe\u0001\u0000"+
		"\u0000\u0000\u00ff\u0100\u0001\u0000\u0000\u0000\u0100\u0101\u0001\u0000"+
		"\u0000\u0000\u0101\u0103\u0005!\u0000\u0000\u0102\u0104\u0003*\u0015\u0000"+
		"\u0103\u0102\u0001\u0000\u0000\u0000\u0103\u0104\u0001\u0000\u0000\u0000"+
		"\u0104\u0105\u0001\u0000\u0000\u0000\u0105\u0106\u0005\u0016\u0000\u0000"+
		"\u0106\u011a\u00030\u0018\u0000\u0107\u0108\u0005\u0011\u0000\u0000\u0108"+
		"\u0109\u0005\u0015\u0000\u0000\u0109\u010a\u0003&\u0013\u0000\u010a\u010b"+
		"\u0005\u0016\u0000\u0000\u010b\u010f\u0005\u0017\u0000\u0000\u010c\u010e"+
		"\u0003,\u0016\u0000\u010d\u010c\u0001\u0000\u0000\u0000\u010e\u0111\u0001"+
		"\u0000\u0000\u0000\u010f\u010d\u0001\u0000\u0000\u0000\u010f\u0110\u0001"+
		"\u0000\u0000\u0000\u0110\u0113\u0001\u0000\u0000\u0000\u0111\u010f\u0001"+
		"\u0000\u0000\u0000\u0112\u0114\u0003.\u0017\u0000\u0113\u0112\u0001\u0000"+
		"\u0000\u0000\u0113\u0114\u0001\u0000\u0000\u0000\u0114\u0115\u0001\u0000"+
		"\u0000\u0000\u0115\u0116\u0005\u0018\u0000\u0000\u0116\u011a\u0001\u0000"+
		"\u0000\u0000\u0117\u011a\u00030\u0018\u0000\u0118\u011a\u0003&\u0013\u0000"+
		"\u0119\u00db\u0001\u0000\u0000\u0000\u0119\u00df\u0001\u0000\u0000\u0000"+
		"\u0119\u00e5\u0001\u0000\u0000\u0000\u0119\u00e7\u0001\u0000\u0000\u0000"+
		"\u0119\u00e8\u0001\u0000\u0000\u0000\u0119\u00e9\u0001\u0000\u0000\u0000"+
		"\u0119\u00f2\u0001\u0000\u0000\u0000\u0119\u00f8\u0001\u0000\u0000\u0000"+
		"\u0119\u0107\u0001\u0000\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000"+
		"\u0119\u0118\u0001\u0000\u0000\u0000\u011a+\u0001\u0000\u0000\u0000\u011b"+
		"\u011c\u0005\u0010\u0000\u0000\u011c\u011d\u0003&\u0013\u0000\u011d\u0124"+
		"\u0005\u001e\u0000\u0000\u011e\u0120\u0003*\u0015\u0000\u011f\u0121\u0005"+
		"!\u0000\u0000\u0120\u011f\u0001\u0000\u0000\u0000\u0120\u0121\u0001\u0000"+
		"\u0000\u0000\u0121\u0123\u0001\u0000\u0000\u0000\u0122\u011e\u0001\u0000"+
		"\u0000\u0000\u0123\u0126\u0001\u0000\u0000\u0000\u0124\u0122\u0001\u0000"+
		"\u0000\u0000\u0124\u0125\u0001\u0000\u0000\u0000\u0125-\u0001\u0000\u0000"+
		"\u0000\u0126\u0124\u0001\u0000\u0000\u0000\u0127\u0128\u0005\u0012\u0000"+
		"\u0000\u0128\u012f\u0005\u001e\u0000\u0000\u0129\u012b\u0003*\u0015\u0000"+
		"\u012a\u012c\u0005!\u0000\u0000\u012b\u012a\u0001\u0000\u0000\u0000\u012b"+
		"\u012c\u0001\u0000\u0000\u0000\u012c\u012e\u0001\u0000\u0000\u0000\u012d"+
		"\u0129\u0001\u0000\u0000\u0000\u012e\u0131\u0001\u0000\u0000\u0000\u012f"+
		"\u012d\u0001\u0000\u0000\u0000\u012f\u0130\u0001\u0000\u0000\u0000\u0130"+
		"/\u0001\u0000\u0000\u0000\u0131\u012f\u0001\u0000\u0000\u0000\u0132\u0139"+
		"\u0005\u0017\u0000\u0000\u0133\u0135\u0003*\u0015\u0000\u0134\u0136\u0005"+
		"!\u0000\u0000\u0135\u0134\u0001\u0000\u0000\u0000\u0135\u0136\u0001\u0000"+
		"\u0000\u0000\u0136\u0138\u0001\u0000\u0000\u0000\u0137\u0133\u0001\u0000"+
		"\u0000\u0000\u0138\u013b\u0001\u0000\u0000\u0000\u0139\u0137\u0001\u0000"+
		"\u0000\u0000\u0139\u013a\u0001\u0000\u0000\u0000\u013a\u013c\u0001\u0000"+
		"\u0000\u0000\u013b\u0139\u0001\u0000\u0000\u0000\u013c\u013d\u0005\u0018"+
		"\u0000\u0000\u013d1\u0001\u0000\u0000\u0000 >MZdjox\u007f\u0087\u008c"+
		"\u0094\u009c\u00a4\u00aa\u00bd\u00d2\u00d4\u00d9\u00e3\u00f0\u00fb\u00ff"+
		"\u0103\u010f\u0113\u0119\u0120\u0124\u012b\u012f\u0135\u0139";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}